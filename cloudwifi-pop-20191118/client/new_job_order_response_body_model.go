// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewJobOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *NewJobOrderResponseBodyData) *NewJobOrderResponseBody
	GetData() *NewJobOrderResponseBodyData
	SetErrorCode(v int32) *NewJobOrderResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *NewJobOrderResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *NewJobOrderResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *NewJobOrderResponseBody
	GetRequestId() *string
}

type NewJobOrderResponseBody struct {
	Data         *NewJobOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                        `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NewJobOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NewJobOrderResponseBody) GoString() string {
	return s.String()
}

func (s *NewJobOrderResponseBody) GetData() *NewJobOrderResponseBodyData {
	return s.Data
}

func (s *NewJobOrderResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *NewJobOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *NewJobOrderResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *NewJobOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NewJobOrderResponseBody) SetData(v *NewJobOrderResponseBodyData) *NewJobOrderResponseBody {
	s.Data = v
	return s
}

func (s *NewJobOrderResponseBody) SetErrorCode(v int32) *NewJobOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NewJobOrderResponseBody) SetErrorMessage(v string) *NewJobOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NewJobOrderResponseBody) SetIsSuccess(v bool) *NewJobOrderResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *NewJobOrderResponseBody) SetRequestId(v string) *NewJobOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *NewJobOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NewJobOrderResponseBodyData struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s NewJobOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s NewJobOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *NewJobOrderResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *NewJobOrderResponseBodyData) SetOrderId(v int64) *NewJobOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *NewJobOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
