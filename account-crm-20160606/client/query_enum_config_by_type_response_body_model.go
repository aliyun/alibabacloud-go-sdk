// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnumConfigByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEnumConfigByTypeResponseBody
	GetCode() *string
	SetData(v *QueryEnumConfigByTypeResponseBodyData) *QueryEnumConfigByTypeResponseBody
	GetData() *QueryEnumConfigByTypeResponseBodyData
	SetMessage(v string) *QueryEnumConfigByTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEnumConfigByTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEnumConfigByTypeResponseBody
	GetSuccess() *bool
}

type QueryEnumConfigByTypeResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryEnumConfigByTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEnumConfigByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEnumConfigByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnumConfigByTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEnumConfigByTypeResponseBody) GetData() *QueryEnumConfigByTypeResponseBodyData {
	return s.Data
}

func (s *QueryEnumConfigByTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEnumConfigByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEnumConfigByTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEnumConfigByTypeResponseBody) SetCode(v string) *QueryEnumConfigByTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBody) SetData(v *QueryEnumConfigByTypeResponseBodyData) *QueryEnumConfigByTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryEnumConfigByTypeResponseBody) SetMessage(v string) *QueryEnumConfigByTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBody) SetRequestId(v string) *QueryEnumConfigByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBody) SetSuccess(v bool) *QueryEnumConfigByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEnumConfigByTypeResponseBodyData struct {
	EnumConfig []*QueryEnumConfigByTypeResponseBodyDataEnumConfig `json:"EnumConfig,omitempty" xml:"EnumConfig,omitempty" type:"Repeated"`
}

func (s QueryEnumConfigByTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEnumConfigByTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEnumConfigByTypeResponseBodyData) GetEnumConfig() []*QueryEnumConfigByTypeResponseBodyDataEnumConfig {
	return s.EnumConfig
}

func (s *QueryEnumConfigByTypeResponseBodyData) SetEnumConfig(v []*QueryEnumConfigByTypeResponseBodyDataEnumConfig) *QueryEnumConfigByTypeResponseBodyData {
	s.EnumConfig = v
	return s
}

func (s *QueryEnumConfigByTypeResponseBodyData) Validate() error {
	if s.EnumConfig != nil {
		for _, item := range s.EnumConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryEnumConfigByTypeResponseBodyDataEnumConfig struct {
	EnumName  *string `json:"enumName,omitempty" xml:"enumName,omitempty"`
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s QueryEnumConfigByTypeResponseBodyDataEnumConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryEnumConfigByTypeResponseBodyDataEnumConfig) GoString() string {
	return s.String()
}

func (s *QueryEnumConfigByTypeResponseBodyDataEnumConfig) GetEnumName() *string {
	return s.EnumName
}

func (s *QueryEnumConfigByTypeResponseBodyDataEnumConfig) GetEnumValue() *string {
	return s.EnumValue
}

func (s *QueryEnumConfigByTypeResponseBodyDataEnumConfig) SetEnumName(v string) *QueryEnumConfigByTypeResponseBodyDataEnumConfig {
	s.EnumName = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBodyDataEnumConfig) SetEnumValue(v string) *QueryEnumConfigByTypeResponseBodyDataEnumConfig {
	s.EnumValue = &v
	return s
}

func (s *QueryEnumConfigByTypeResponseBodyDataEnumConfig) Validate() error {
	return dara.Validate(s)
}
