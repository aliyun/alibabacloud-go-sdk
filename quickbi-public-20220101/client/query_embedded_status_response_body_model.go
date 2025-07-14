// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmbeddedStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryEmbeddedStatusResponseBody
	GetRequestId() *string
	SetResult(v bool) *QueryEmbeddedStatusResponseBody
	GetResult() *bool
	SetSuccess(v bool) *QueryEmbeddedStatusResponseBody
	GetSuccess() *bool
}

type QueryEmbeddedStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the work is enabled for embedding. Valid values:
	//
	// 	- true: embedded
	//
	// 	- false: not embedded
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

func (s QueryEmbeddedStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEmbeddedStatusResponseBody) GetResult() *bool {
	return s.Result
}

func (s *QueryEmbeddedStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEmbeddedStatusResponseBody) SetRequestId(v string) *QueryEmbeddedStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmbeddedStatusResponseBody) SetResult(v bool) *QueryEmbeddedStatusResponseBody {
	s.Result = &v
	return s
}

func (s *QueryEmbeddedStatusResponseBody) SetSuccess(v bool) *QueryEmbeddedStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEmbeddedStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
