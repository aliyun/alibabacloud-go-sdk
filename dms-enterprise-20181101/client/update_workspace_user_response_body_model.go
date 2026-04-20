// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateWorkspaceUserResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateWorkspaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkspaceUserResponseBody
	GetSuccess() *bool
}

type UpdateWorkspaceUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C5B8E84B-42B6-4374-AD5A-6264E1753325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateWorkspaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkspaceUserResponseBody) SetData(v bool) *UpdateWorkspaceUserResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWorkspaceUserResponseBody) SetRequestId(v string) *UpdateWorkspaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceUserResponseBody) SetSuccess(v bool) *UpdateWorkspaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
