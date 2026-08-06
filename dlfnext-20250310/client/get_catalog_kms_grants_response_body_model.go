// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogKmsGrantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataAccessRoleArn(v string) *GetCatalogKmsGrantsResponseBody
	GetDataAccessRoleArn() *string
	SetKeyPolicyStatement(v string) *GetCatalogKmsGrantsResponseBody
	GetKeyPolicyStatement() *string
	SetRegion(v string) *GetCatalogKmsGrantsResponseBody
	GetRegion() *string
	SetWorkflowRoleArn(v string) *GetCatalogKmsGrantsResponseBody
	GetWorkflowRoleArn() *string
}

type GetCatalogKmsGrantsResponseBody struct {
	// The ARN of the RAM role used by DLF to access catalog data. When configuring the KMS key policy, you must grant this role permissions to use the customer master key.
	//
	// example:
	//
	// acs:ram::123456789012****:role/AliyunDlfNextDataAccessRole
	DataAccessRoleArn *string `json:"dataAccessRoleArn,omitempty" xml:"dataAccessRoleArn,omitempty"`
	// The authorization statement that must be added to the customer master key policy. This statement grants the DLF data access role corresponding to dataAccessRoleArn the KMS permissions required for data encryption and decryption.
	//
	// example:
	//
	// {"Sid":"AllowDLFDataAccess","Effect":"Allow","Principal":{"RAM":["acs:ram::123456789012****:role/
	//
	//   AliyunDlfNextDataAccessRole"]},"Action":["kms:Decrypt","kms:GenerateDataKey"],"Resource":["*"]}
	KeyPolicyStatement *string `json:"keyPolicyStatement,omitempty" xml:"keyPolicyStatement,omitempty"`
	// The region ID to which the catalog belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The DLF workflow role ARN. In the current version, the workflow role is not granted customer master key access permissions based on the least privilege principle. Therefore, this field returns an empty value.
	//
	// example:
	//
	// null
	WorkflowRoleArn *string `json:"workflowRoleArn,omitempty" xml:"workflowRoleArn,omitempty"`
}

func (s GetCatalogKmsGrantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogKmsGrantsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogKmsGrantsResponseBody) GetDataAccessRoleArn() *string {
	return s.DataAccessRoleArn
}

func (s *GetCatalogKmsGrantsResponseBody) GetKeyPolicyStatement() *string {
	return s.KeyPolicyStatement
}

func (s *GetCatalogKmsGrantsResponseBody) GetRegion() *string {
	return s.Region
}

func (s *GetCatalogKmsGrantsResponseBody) GetWorkflowRoleArn() *string {
	return s.WorkflowRoleArn
}

func (s *GetCatalogKmsGrantsResponseBody) SetDataAccessRoleArn(v string) *GetCatalogKmsGrantsResponseBody {
	s.DataAccessRoleArn = &v
	return s
}

func (s *GetCatalogKmsGrantsResponseBody) SetKeyPolicyStatement(v string) *GetCatalogKmsGrantsResponseBody {
	s.KeyPolicyStatement = &v
	return s
}

func (s *GetCatalogKmsGrantsResponseBody) SetRegion(v string) *GetCatalogKmsGrantsResponseBody {
	s.Region = &v
	return s
}

func (s *GetCatalogKmsGrantsResponseBody) SetWorkflowRoleArn(v string) *GetCatalogKmsGrantsResponseBody {
	s.WorkflowRoleArn = &v
	return s
}

func (s *GetCatalogKmsGrantsResponseBody) Validate() error {
	return dara.Validate(s)
}
