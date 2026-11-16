// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddProjectMembersResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddProjectMembersResponseBody
	GetResult() *bool
}

type AddProjectMembersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01A1E2B8-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddProjectMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddProjectMembersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddProjectMembersResponseBody) SetRequestId(v string) *AddProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectMembersResponseBody) SetResult(v bool) *AddProjectMembersResponseBody {
	s.Result = &v
	return s
}

func (s *AddProjectMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
