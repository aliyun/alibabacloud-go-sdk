// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataTrackOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateDataTrackOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateDataTrackOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataTrackOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataTrackOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataTrackOrderResponseBody
	GetSuccess() *bool
}

type CreateDataTrackOrderResponseBody struct {
	// The IDs of the data tracking tickets.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataTrackOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataTrackOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataTrackOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataTrackOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataTrackOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataTrackOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataTrackOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataTrackOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataTrackOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataTrackOrderResponseBody) SetErrorCode(v string) *CreateDataTrackOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataTrackOrderResponseBody) SetErrorMessage(v string) *CreateDataTrackOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataTrackOrderResponseBody) SetRequestId(v string) *CreateDataTrackOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataTrackOrderResponseBody) SetSuccess(v bool) *CreateDataTrackOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataTrackOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
