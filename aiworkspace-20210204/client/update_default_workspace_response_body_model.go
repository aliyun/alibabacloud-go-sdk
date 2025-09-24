// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDefaultWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDefaultWorkspaceResponseBody
	GetRequestId() *string
}

type UpdateDefaultWorkspaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 17915******4216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDefaultWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDefaultWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDefaultWorkspaceResponseBody) SetRequestId(v string) *UpdateDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDefaultWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
