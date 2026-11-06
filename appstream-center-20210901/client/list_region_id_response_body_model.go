// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRegionIdResponseBody
	GetCode() *string
	SetCount(v int32) *ListRegionIdResponseBody
	GetCount() *int32
	SetData(v []*string) *ListRegionIdResponseBody
	GetData() []*string
	SetMessage(v string) *ListRegionIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRegionIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRegionIdResponseBody
	GetSuccess() *bool
}

type ListRegionIdResponseBody struct {
	// The error code. This parameter is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of records returned in this list.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of region IDs available for image services under the selected product for the current account.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error or result message. This parameter describes the cause when a request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRegionIdResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListRegionIdResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListRegionIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRegionIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRegionIdResponseBody) SetCode(v string) *ListRegionIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionIdResponseBody) SetCount(v int32) *ListRegionIdResponseBody {
	s.Count = &v
	return s
}

func (s *ListRegionIdResponseBody) SetData(v []*string) *ListRegionIdResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionIdResponseBody) SetMessage(v string) *ListRegionIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionIdResponseBody) SetRequestId(v string) *ListRegionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionIdResponseBody) SetSuccess(v bool) *ListRegionIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionIdResponseBody) Validate() error {
	return dara.Validate(s)
}
