// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteResourceResponseBody
	GetSuccess() *bool
}

type DeleteResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetSuccess(v bool) *DeleteResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
