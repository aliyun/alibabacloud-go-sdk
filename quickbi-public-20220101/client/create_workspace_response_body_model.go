// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateWorkspaceResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateWorkspaceResponseBody
	GetSuccess() *bool
}

type CreateWorkspaceResponseBody struct {
	// example:
	//
	// 685072****************4e79e718f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12423twfasva********
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetResult(v string) *CreateWorkspaceResponseBody {
	s.Result = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
