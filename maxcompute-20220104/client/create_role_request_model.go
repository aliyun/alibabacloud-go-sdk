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
	// {"name": "role_name","type": "resource/admin","policy": "", //policy content. Not required if the type is acl."acl": { // Not required if the type is policy."table": [{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"resource":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"function":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"package":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"project":[{"name": "", "actions":[]}], //Only the current project is available on the console."instance":[{"name": "", "actions":[]}] //Only 	- is supported for the name. }}// 	- is supported for the name.
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
