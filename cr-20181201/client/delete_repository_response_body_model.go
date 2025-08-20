// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteRepositoryResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteRepositoryResponseBody
	GetRequestId() *string
}

type DeleteRepositoryResponseBody struct {
	// Return values
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 47DD9D56-09A0-4C52-B520-C3805DBAB96B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepositoryResponseBody) SetCode(v string) *DeleteRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetIsSuccess(v bool) *DeleteRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
