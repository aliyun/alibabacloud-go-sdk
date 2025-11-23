// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteWorkspaceResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkspaceResponseBody
	GetSuccess() *bool
}

type DeleteWorkspaceResponseBody struct {
	// Indicates whether the workspace is deleted successfully.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is called successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkspaceResponseBody) SetData(v string) *DeleteWorkspaceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetSuccess(v bool) *DeleteWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
