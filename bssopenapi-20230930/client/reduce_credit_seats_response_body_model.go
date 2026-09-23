// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceCreditSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReduceCreditSeatsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ReduceCreditSeatsResponseBody
	GetCode() *string
	SetData(v *ReduceCreditSeatsResponseBodyData) *ReduceCreditSeatsResponseBody
	GetData() *ReduceCreditSeatsResponseBodyData
	SetMessage(v string) *ReduceCreditSeatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReduceCreditSeatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReduceCreditSeatsResponseBody
	GetSuccess() *bool
}

type ReduceCreditSeatsResponseBody struct {
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ReduceCreditSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReduceCreditSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReduceCreditSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *ReduceCreditSeatsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReduceCreditSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReduceCreditSeatsResponseBody) GetData() *ReduceCreditSeatsResponseBodyData {
	return s.Data
}

func (s *ReduceCreditSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReduceCreditSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReduceCreditSeatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReduceCreditSeatsResponseBody) SetAccessDeniedDetail(v string) *ReduceCreditSeatsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReduceCreditSeatsResponseBody) SetCode(v string) *ReduceCreditSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *ReduceCreditSeatsResponseBody) SetData(v *ReduceCreditSeatsResponseBodyData) *ReduceCreditSeatsResponseBody {
	s.Data = v
	return s
}

func (s *ReduceCreditSeatsResponseBody) SetMessage(v string) *ReduceCreditSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *ReduceCreditSeatsResponseBody) SetRequestId(v string) *ReduceCreditSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReduceCreditSeatsResponseBody) SetSuccess(v bool) *ReduceCreditSeatsResponseBody {
	s.Success = &v
	return s
}

func (s *ReduceCreditSeatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReduceCreditSeatsResponseBodyData struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo  *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReduceCreditSeatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReduceCreditSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReduceCreditSeatsResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReduceCreditSeatsResponseBodyData) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *ReduceCreditSeatsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReduceCreditSeatsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ReduceCreditSeatsResponseBodyData) SetErrorCode(v string) *ReduceCreditSeatsResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *ReduceCreditSeatsResponseBodyData) SetErrorInfo(v string) *ReduceCreditSeatsResponseBodyData {
	s.ErrorInfo = &v
	return s
}

func (s *ReduceCreditSeatsResponseBodyData) SetInstanceId(v string) *ReduceCreditSeatsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ReduceCreditSeatsResponseBodyData) SetSuccess(v bool) *ReduceCreditSeatsResponseBodyData {
	s.Success = &v
	return s
}

func (s *ReduceCreditSeatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
