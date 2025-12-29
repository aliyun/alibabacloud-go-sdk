// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeRoleTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNodeRoleTagResponseBody
	GetRequestId() *string
}

type CreateNodeRoleTagResponseBody struct {
	// example:
	//
	// B0B7DE3E-xxxx-xxxx-xxxx-0B16F8834E0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeRoleTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeRoleTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeRoleTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeRoleTagResponseBody) SetRequestId(v string) *CreateNodeRoleTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeRoleTagResponseBody) Validate() error {
	return dara.Validate(s)
}
