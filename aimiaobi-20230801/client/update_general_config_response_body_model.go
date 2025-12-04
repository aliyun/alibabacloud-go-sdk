// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneralConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGeneralConfigResponseBody
	GetCode() *string
	SetData(v *UpdateGeneralConfigResponseBodyData) *UpdateGeneralConfigResponseBody
	GetData() *UpdateGeneralConfigResponseBodyData
	SetHttpStatusCode(v int32) *UpdateGeneralConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGeneralConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGeneralConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGeneralConfigResponseBody
	GetSuccess() *bool
}

type UpdateGeneralConfigResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateGeneralConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGeneralConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneralConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGeneralConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGeneralConfigResponseBody) GetData() *UpdateGeneralConfigResponseBodyData {
	return s.Data
}

func (s *UpdateGeneralConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGeneralConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGeneralConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGeneralConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGeneralConfigResponseBody) SetCode(v string) *UpdateGeneralConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGeneralConfigResponseBody) SetData(v *UpdateGeneralConfigResponseBodyData) *UpdateGeneralConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGeneralConfigResponseBody) SetHttpStatusCode(v int32) *UpdateGeneralConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGeneralConfigResponseBody) SetMessage(v string) *UpdateGeneralConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGeneralConfigResponseBody) SetRequestId(v string) *UpdateGeneralConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGeneralConfigResponseBody) SetSuccess(v bool) *UpdateGeneralConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGeneralConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGeneralConfigResponseBodyData struct {
	// example:
	//
	// xxx
	ConfigDesc *string `json:"ConfigDesc,omitempty" xml:"ConfigDesc,omitempty"`
	// example:
	//
	// xx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// xxx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// example:
	//
	// xx
	ConfigValueType *string `json:"ConfigValueType,omitempty" xml:"ConfigValueType,omitempty"`
}

func (s UpdateGeneralConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneralConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGeneralConfigResponseBodyData) GetConfigDesc() *string {
	return s.ConfigDesc
}

func (s *UpdateGeneralConfigResponseBodyData) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *UpdateGeneralConfigResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateGeneralConfigResponseBodyData) GetConfigValueType() *string {
	return s.ConfigValueType
}

func (s *UpdateGeneralConfigResponseBodyData) SetConfigDesc(v string) *UpdateGeneralConfigResponseBodyData {
	s.ConfigDesc = &v
	return s
}

func (s *UpdateGeneralConfigResponseBodyData) SetConfigKey(v string) *UpdateGeneralConfigResponseBodyData {
	s.ConfigKey = &v
	return s
}

func (s *UpdateGeneralConfigResponseBodyData) SetConfigValue(v string) *UpdateGeneralConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *UpdateGeneralConfigResponseBodyData) SetConfigValueType(v string) *UpdateGeneralConfigResponseBodyData {
	s.ConfigValueType = &v
	return s
}

func (s *UpdateGeneralConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
