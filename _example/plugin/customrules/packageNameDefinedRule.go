package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// PackageNameDefinedRule verifies that the package name defined.
type PackageNameDefinedRule struct{}

// NewPackageNameDefinedRule creates a new PackageNameDefinedRule.
func NewPackageNameDefinedRule() PackageNameDefinedRule {
	return PackageNameDefinedRule{}
}

// ID returns the ID of this rule.
func (r PackageNameDefinedRule) ID() string {
	return "PACKAGE_NAME_DEFINED"
}

// Purpose returns the purpose of this rule.
func (r PackageNameDefinedRule) Purpose() string {
	return "Verifies that the package name is defined."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r PackageNameDefinedRule) IsOfficial() bool {
	return false
}

// Apply applies the rule to the proto.
func (r PackageNameDefinedRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &packageNameDefinedVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type packageNameDefinedVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitPackage checks the package.
func (v *packageNameDefinedVisitor) VisitPackage(p *parser.Package) bool {
	name := p.Name
	if !isPackageDefined(name) {
		v.AddFailuref(p.Meta.Pos, "Package name is not defined.")
	}
	return false
}

func isPackageDefined(packageName string) bool {
	print(packageName)
	return len(packageName) > 0
}
