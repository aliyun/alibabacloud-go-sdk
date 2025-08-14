// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSourceResponseBody
	GetSuccess() *bool
}

type DeleteSourceResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
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

func (s DeleteSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSourceResponseBody) SetRequestId(v string) *DeleteSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceResponseBody) SetSuccess(v bool) *DeleteSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
