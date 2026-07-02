// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAuthorizationByUserIdResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteAuthorizationByUserIdResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteAuthorizationByUserIdResponseBody
	GetSuccess() *bool
}

type DeleteAuthorizationByUserIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46e53***********70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAuthorizationByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuthorizationByUserIdResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteAuthorizationByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAuthorizationByUserIdResponseBody) SetRequestId(v string) *DeleteAuthorizationByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuthorizationByUserIdResponseBody) SetResult(v bool) *DeleteAuthorizationByUserIdResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteAuthorizationByUserIdResponseBody) SetSuccess(v bool) *DeleteAuthorizationByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAuthorizationByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
