// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSourceLocationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSourceLocationResponseBody
	GetSuccess() *bool
}

type DeleteSourceLocationResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid value:
	//
	// 	- true: The request succeeded.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSourceLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceLocationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSourceLocationResponseBody) SetRequestId(v string) *DeleteSourceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceLocationResponseBody) SetSuccess(v bool) *DeleteSourceLocationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSourceLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
