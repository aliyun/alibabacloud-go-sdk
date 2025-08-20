// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRepoTagResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteRepoTagResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteRepoTagResponseBody
	GetRequestId() *string
}

type DeleteRepoTagResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 96E66B3A-C81A-48BE-ACD6-C0AB1F9313C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRepoTagResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteRepoTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepoTagResponseBody) SetCode(v string) *DeleteRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoTagResponseBody) SetIsSuccess(v bool) *DeleteRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoTagResponseBody) SetRequestId(v string) *DeleteRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepoTagResponseBody) Validate() error {
	return dara.Validate(s)
}
