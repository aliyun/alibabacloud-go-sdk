// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostPaidBindRelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePostPaidBindRelResponseBody
	GetCode() *string
	SetData(v *UpdatePostPaidBindRelResponseBodyData) *UpdatePostPaidBindRelResponseBody
	GetData() *UpdatePostPaidBindRelResponseBodyData
	SetMessage(v string) *UpdatePostPaidBindRelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePostPaidBindRelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePostPaidBindRelResponseBody
	GetSuccess() *bool
}

type UpdatePostPaidBindRelResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdatePostPaidBindRelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePostPaidBindRelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePostPaidBindRelResponseBody) GetData() *UpdatePostPaidBindRelResponseBodyData {
	return s.Data
}

func (s *UpdatePostPaidBindRelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePostPaidBindRelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostPaidBindRelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePostPaidBindRelResponseBody) SetCode(v string) *UpdatePostPaidBindRelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBody) SetData(v *UpdatePostPaidBindRelResponseBodyData) *UpdatePostPaidBindRelResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePostPaidBindRelResponseBody) SetMessage(v string) *UpdatePostPaidBindRelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBody) SetRequestId(v string) *UpdatePostPaidBindRelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBody) SetSuccess(v bool) *UpdatePostPaidBindRelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePostPaidBindRelResponseBodyData struct {
	Body *UpdatePostPaidBindRelResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s UpdatePostPaidBindRelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelResponseBodyData) GetBody() *UpdatePostPaidBindRelResponseBodyDataBody {
	return s.Body
}

func (s *UpdatePostPaidBindRelResponseBodyData) SetBody(v *UpdatePostPaidBindRelResponseBodyDataBody) *UpdatePostPaidBindRelResponseBodyData {
	s.Body = v
	return s
}

func (s *UpdatePostPaidBindRelResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePostPaidBindRelResponseBodyDataBody struct {
	BindCount  *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode *int32  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s UpdatePostPaidBindRelResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) GetBindCount() *int64 {
	return s.BindCount
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) SetBindCount(v int64) *UpdatePostPaidBindRelResponseBodyDataBody {
	s.BindCount = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) SetRequestId(v string) *UpdatePostPaidBindRelResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) SetResultCode(v int32) *UpdatePostPaidBindRelResponseBodyDataBody {
	s.ResultCode = &v
	return s
}

func (s *UpdatePostPaidBindRelResponseBodyDataBody) Validate() error {
	return dara.Validate(s)
}
