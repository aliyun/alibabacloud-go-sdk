// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateWorkspaceResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateWorkspaceResponseBody
	GetSuccess() *string
}

type UpdateWorkspaceResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// C51420E3-144A-4A94-B473-8662FCF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateWorkspaceResponseBody) SetData(v string) *UpdateWorkspaceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetSuccess(v string) *UpdateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
