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
	DataAccessRoleArn  *string `json:"dataAccessRoleArn,omitempty" xml:"dataAccessRoleArn,omitempty"`
	KeyPolicyStatement *string `json:"keyPolicyStatement,omitempty" xml:"keyPolicyStatement,omitempty"`
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	WorkflowRoleArn    *string `json:"workflowRoleArn,omitempty" xml:"workflowRoleArn,omitempty"`
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
