// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserTagMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserTagMetaResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserTagMetaResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserTagMetaResponseBody
	GetSuccess() *bool
}

type DeleteUserTagMetaResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the deleted tag is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserTagMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserTagMetaResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserTagMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserTagMetaResponseBody) SetRequestId(v string) *DeleteUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserTagMetaResponseBody) SetResult(v bool) *DeleteUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserTagMetaResponseBody) SetSuccess(v bool) *DeleteUserTagMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserTagMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
