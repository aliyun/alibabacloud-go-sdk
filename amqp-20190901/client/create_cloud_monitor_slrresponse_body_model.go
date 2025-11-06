// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMonitorSLRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateCloudMonitorSLRResponseBody
	GetCode() *int32
	SetData(v bool) *CreateCloudMonitorSLRResponseBody
	GetData() *bool
	SetMessage(v string) *CreateCloudMonitorSLRResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCloudMonitorSLRResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudMonitorSLRResponseBody
	GetSuccess() *bool
}

type CreateCloudMonitorSLRResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCloudMonitorSLRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMonitorSLRResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudMonitorSLRResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateCloudMonitorSLRResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateCloudMonitorSLRResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCloudMonitorSLRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudMonitorSLRResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudMonitorSLRResponseBody) SetCode(v int32) *CreateCloudMonitorSLRResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCloudMonitorSLRResponseBody) SetData(v bool) *CreateCloudMonitorSLRResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCloudMonitorSLRResponseBody) SetMessage(v string) *CreateCloudMonitorSLRResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCloudMonitorSLRResponseBody) SetRequestId(v string) *CreateCloudMonitorSLRResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudMonitorSLRResponseBody) SetSuccess(v bool) *CreateCloudMonitorSLRResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudMonitorSLRResponseBody) Validate() error {
	return dara.Validate(s)
}
