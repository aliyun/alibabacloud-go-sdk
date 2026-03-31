// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateRoleRequest
	GetBody() *string
}

type CreateRoleRequest struct {
	// The request body parameters. For valid values, see [MaxCompute permissions](https://help.aliyun.com/document_detail/27935.html).
	//
	// example:
	//
	// {"name": "role_name","type": "resource/adminn","policy": "", // The document of the policy. This parameter is not required if an access-control list (ACL) is used. "acl": { // This parameter is not required if a policy is used. "table": [{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"resource":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"function":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"package":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"project":[{"name": "", "actions":[]}], // Only the current project is displayed in the console. "instance":[{"name": "", "actions":[]}] // The parameter name must be set to an asterisk (\\*) in the console. }}// An asterisk (\\*) can be specified for name.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetBody() *string {
	return s.Body
}

func (s *CreateRoleRequest) SetBody(v string) *CreateRoleRequest {
	s.Body = &v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	return dara.Validate(s)
}
