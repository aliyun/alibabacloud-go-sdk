// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteChainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteChainResponseBody
	GetRequestId() *string
}

type DeleteChainResponseBody struct {
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
	// DB1809A8-E1C8-5707-BAF8-D4FC1C11****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChainResponseBody) SetCode(v string) *DeleteChainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChainResponseBody) SetIsSuccess(v bool) *DeleteChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChainResponseBody) SetRequestId(v string) *DeleteChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChainResponseBody) Validate() error {
	return dara.Validate(s)
}
