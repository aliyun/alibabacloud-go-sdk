// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalitySettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateLocalitySettingResponseBody
	GetCode() *int32
	SetData(v *UpdateLocalitySettingResponseBodyData) *UpdateLocalitySettingResponseBody
	GetData() *UpdateLocalitySettingResponseBodyData
	SetHttpStatusCode(v int32) *UpdateLocalitySettingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLocalitySettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLocalitySettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLocalitySettingResponseBody
	GetSuccess() *bool
}

type UpdateLocalitySettingResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateLocalitySettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLocalitySettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalitySettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocalitySettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateLocalitySettingResponseBody) GetData() *UpdateLocalitySettingResponseBodyData {
	return s.Data
}

func (s *UpdateLocalitySettingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLocalitySettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLocalitySettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLocalitySettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLocalitySettingResponseBody) SetCode(v int32) *UpdateLocalitySettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLocalitySettingResponseBody) SetData(v *UpdateLocalitySettingResponseBodyData) *UpdateLocalitySettingResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLocalitySettingResponseBody) SetHttpStatusCode(v int32) *UpdateLocalitySettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLocalitySettingResponseBody) SetMessage(v string) *UpdateLocalitySettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLocalitySettingResponseBody) SetRequestId(v string) *UpdateLocalitySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLocalitySettingResponseBody) SetSuccess(v bool) *UpdateLocalitySettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLocalitySettingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLocalitySettingResponseBodyData struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 15
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLocalitySettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalitySettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLocalitySettingResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateLocalitySettingResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateLocalitySettingResponseBodyData) SetEnabled(v bool) *UpdateLocalitySettingResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdateLocalitySettingResponseBodyData) SetThreshold(v float32) *UpdateLocalitySettingResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateLocalitySettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
