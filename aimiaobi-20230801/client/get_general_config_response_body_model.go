// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneralConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGeneralConfigResponseBody
	GetCode() *string
	SetData(v *GetGeneralConfigResponseBodyData) *GetGeneralConfigResponseBody
	GetData() *GetGeneralConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetGeneralConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGeneralConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGeneralConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGeneralConfigResponseBody
	GetSuccess() *bool
}

type GetGeneralConfigResponseBody struct {
	// example:
	//
	// NoData
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetGeneralConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetGeneralConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGeneralConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneralConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGeneralConfigResponseBody) GetData() *GetGeneralConfigResponseBodyData {
	return s.Data
}

func (s *GetGeneralConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGeneralConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGeneralConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGeneralConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGeneralConfigResponseBody) SetCode(v string) *GetGeneralConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetGeneralConfigResponseBody) SetData(v *GetGeneralConfigResponseBodyData) *GetGeneralConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetGeneralConfigResponseBody) SetHttpStatusCode(v int32) *GetGeneralConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGeneralConfigResponseBody) SetMessage(v string) *GetGeneralConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetGeneralConfigResponseBody) SetRequestId(v string) *GetGeneralConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGeneralConfigResponseBody) SetSuccess(v bool) *GetGeneralConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetGeneralConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGeneralConfigResponseBodyData struct {
	// example:
	//
	// xx
	ConfigDesc *string `json:"ConfigDesc,omitempty" xml:"ConfigDesc,omitempty"`
	// example:
	//
	// xx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// xx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// example:
	//
	// xx
	ConfigValueType *string `json:"ConfigValueType,omitempty" xml:"ConfigValueType,omitempty"`
}

func (s GetGeneralConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGeneralConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGeneralConfigResponseBodyData) GetConfigDesc() *string {
	return s.ConfigDesc
}

func (s *GetGeneralConfigResponseBodyData) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *GetGeneralConfigResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetGeneralConfigResponseBodyData) GetConfigValueType() *string {
	return s.ConfigValueType
}

func (s *GetGeneralConfigResponseBodyData) SetConfigDesc(v string) *GetGeneralConfigResponseBodyData {
	s.ConfigDesc = &v
	return s
}

func (s *GetGeneralConfigResponseBodyData) SetConfigKey(v string) *GetGeneralConfigResponseBodyData {
	s.ConfigKey = &v
	return s
}

func (s *GetGeneralConfigResponseBodyData) SetConfigValue(v string) *GetGeneralConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *GetGeneralConfigResponseBodyData) SetConfigValueType(v string) *GetGeneralConfigResponseBodyData {
	s.ConfigValueType = &v
	return s
}

func (s *GetGeneralConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
