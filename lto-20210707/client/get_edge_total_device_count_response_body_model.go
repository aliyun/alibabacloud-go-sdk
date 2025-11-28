// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTotalDeviceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEdgeTotalDeviceCountResponseBody
	GetCode() *string
	SetData(v int64) *GetEdgeTotalDeviceCountResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *GetEdgeTotalDeviceCountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetEdgeTotalDeviceCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEdgeTotalDeviceCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEdgeTotalDeviceCountResponseBody
	GetSuccess() *bool
}

type GetEdgeTotalDeviceCountResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEdgeTotalDeviceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTotalDeviceCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeTotalDeviceCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetCode(v string) *GetEdgeTotalDeviceCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetData(v int64) *GetEdgeTotalDeviceCountResponseBody {
	s.Data = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetHttpStatusCode(v int32) *GetEdgeTotalDeviceCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetMessage(v string) *GetEdgeTotalDeviceCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetRequestId(v string) *GetEdgeTotalDeviceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) SetSuccess(v bool) *GetEdgeTotalDeviceCountResponseBody {
	s.Success = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponseBody) Validate() error {
	return dara.Validate(s)
}
