// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSmartqStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDatasetSmartqStatusResponseBody
	GetRequestId() *string
	SetResult(v bool) *QueryDatasetSmartqStatusResponseBody
	GetResult() *bool
	SetSuccess(v bool) *QueryDatasetSmartqStatusResponseBody
	GetSuccess() *bool
}

type QueryDatasetSmartqStatusResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result of the API execution. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
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

func (s QueryDatasetSmartqStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSmartqStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetSmartqStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetSmartqStatusResponseBody) GetResult() *bool {
	return s.Result
}

func (s *QueryDatasetSmartqStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDatasetSmartqStatusResponseBody) SetRequestId(v string) *QueryDatasetSmartqStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetSmartqStatusResponseBody) SetResult(v bool) *QueryDatasetSmartqStatusResponseBody {
	s.Result = &v
	return s
}

func (s *QueryDatasetSmartqStatusResponseBody) SetSuccess(v bool) *QueryDatasetSmartqStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDatasetSmartqStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
