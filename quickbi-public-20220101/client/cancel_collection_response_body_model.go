// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCollectionResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelCollectionResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CancelCollectionResponseBody
	GetSuccess() *bool
}

type CancelCollectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
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

func (s CancelCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCollectionResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelCollectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelCollectionResponseBody) SetRequestId(v string) *CancelCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCollectionResponseBody) SetResult(v bool) *CancelCollectionResponseBody {
	s.Result = &v
	return s
}

func (s *CancelCollectionResponseBody) SetSuccess(v bool) *CancelCollectionResponseBody {
	s.Success = &v
	return s
}

func (s *CancelCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
