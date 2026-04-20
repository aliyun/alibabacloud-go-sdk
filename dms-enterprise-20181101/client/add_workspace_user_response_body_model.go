// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddWorkspaceUserResponseBody
	GetData() *bool
	SetRequestId(v string) *AddWorkspaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddWorkspaceUserResponseBody
	GetSuccess() *bool
}

type AddWorkspaceUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B4B07137-F6AE-4756-8474-7F92BB6C4E04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorkspaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddWorkspaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkspaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWorkspaceUserResponseBody) SetData(v bool) *AddWorkspaceUserResponseBody {
	s.Data = &v
	return s
}

func (s *AddWorkspaceUserResponseBody) SetRequestId(v string) *AddWorkspaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceUserResponseBody) SetSuccess(v bool) *AddWorkspaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *AddWorkspaceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
