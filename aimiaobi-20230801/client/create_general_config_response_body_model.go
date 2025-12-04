// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneralConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGeneralConfigResponseBody
	GetCode() *string
	SetData(v *CreateGeneralConfigResponseBodyData) *CreateGeneralConfigResponseBody
	GetData() *CreateGeneralConfigResponseBodyData
	SetHttpStatusCode(v int32) *CreateGeneralConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateGeneralConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGeneralConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGeneralConfigResponseBody
	GetSuccess() *bool
}

type CreateGeneralConfigResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateGeneralConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGeneralConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneralConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGeneralConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGeneralConfigResponseBody) GetData() *CreateGeneralConfigResponseBodyData {
	return s.Data
}

func (s *CreateGeneralConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGeneralConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGeneralConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGeneralConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGeneralConfigResponseBody) SetCode(v string) *CreateGeneralConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGeneralConfigResponseBody) SetData(v *CreateGeneralConfigResponseBodyData) *CreateGeneralConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateGeneralConfigResponseBody) SetHttpStatusCode(v int32) *CreateGeneralConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGeneralConfigResponseBody) SetMessage(v string) *CreateGeneralConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGeneralConfigResponseBody) SetRequestId(v string) *CreateGeneralConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGeneralConfigResponseBody) SetSuccess(v bool) *CreateGeneralConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGeneralConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGeneralConfigResponseBodyData struct {
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

func (s CreateGeneralConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneralConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGeneralConfigResponseBodyData) GetConfigDesc() *string {
	return s.ConfigDesc
}

func (s *CreateGeneralConfigResponseBodyData) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *CreateGeneralConfigResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *CreateGeneralConfigResponseBodyData) GetConfigValueType() *string {
	return s.ConfigValueType
}

func (s *CreateGeneralConfigResponseBodyData) SetConfigDesc(v string) *CreateGeneralConfigResponseBodyData {
	s.ConfigDesc = &v
	return s
}

func (s *CreateGeneralConfigResponseBodyData) SetConfigKey(v string) *CreateGeneralConfigResponseBodyData {
	s.ConfigKey = &v
	return s
}

func (s *CreateGeneralConfigResponseBodyData) SetConfigValue(v string) *CreateGeneralConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *CreateGeneralConfigResponseBodyData) SetConfigValueType(v string) *CreateGeneralConfigResponseBodyData {
	s.ConfigValueType = &v
	return s
}

func (s *CreateGeneralConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
