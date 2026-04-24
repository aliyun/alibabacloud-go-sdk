// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateApiKeyResponseBodyData) *CreateApiKeyResponseBody
	GetData() *CreateApiKeyResponseBodyData
	SetMessage(v string) *CreateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApiKeyResponseBody
	GetSuccess() *bool
}

type CreateApiKeyResponseBody struct {
	Data *CreateApiKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBody) GetData() *CreateApiKeyResponseBodyData {
	return s.Data
}

func (s *CreateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApiKeyResponseBody) SetData(v *CreateApiKeyResponseBodyData) *CreateApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *CreateApiKeyResponseBody) SetMessage(v string) *CreateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetRequestId(v string) *CreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetSuccess(v bool) *CreateApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApiKeyResponseBodyData struct {
	// example:
	//
	// http://xxx.yy/v1
	BaseUrl       *string                                      `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	CustomKeyList []*CreateApiKeyResponseBodyDataCustomKeyList `json:"CustomKeyList,omitempty" xml:"CustomKeyList,omitempty" type:"Repeated"`
	// example:
	//
	// sk-rds-xxx
	SystemApiKey *string `json:"SystemApiKey,omitempty" xml:"SystemApiKey,omitempty"`
}

func (s CreateApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateApiKeyResponseBodyData) GetCustomKeyList() []*CreateApiKeyResponseBodyDataCustomKeyList {
	return s.CustomKeyList
}

func (s *CreateApiKeyResponseBodyData) GetSystemApiKey() *string {
	return s.SystemApiKey
}

func (s *CreateApiKeyResponseBodyData) SetBaseUrl(v string) *CreateApiKeyResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *CreateApiKeyResponseBodyData) SetCustomKeyList(v []*CreateApiKeyResponseBodyDataCustomKeyList) *CreateApiKeyResponseBodyData {
	s.CustomKeyList = v
	return s
}

func (s *CreateApiKeyResponseBodyData) SetSystemApiKey(v string) *CreateApiKeyResponseBodyData {
	s.SystemApiKey = &v
	return s
}

func (s *CreateApiKeyResponseBodyData) Validate() error {
	if s.CustomKeyList != nil {
		for _, item := range s.CustomKeyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApiKeyResponseBodyDataCustomKeyList struct {
	// Api Key
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// 0.2
	LimitRate *float32 `json:"LimitRate,omitempty" xml:"LimitRate,omitempty"`
	// example:
	//
	// fixed
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// example:
	//
	// 100000
	TokenQuota *int64 `json:"TokenQuota,omitempty" xml:"TokenQuota,omitempty"`
}

func (s CreateApiKeyResponseBodyDataCustomKeyList) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBodyDataCustomKeyList) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) GetKeyName() *string {
	return s.KeyName
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) GetLimitRate() *float32 {
	return s.LimitRate
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) GetTokenQuota() *int64 {
	return s.TokenQuota
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) SetApiKey(v string) *CreateApiKeyResponseBodyDataCustomKeyList {
	s.ApiKey = &v
	return s
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) SetKeyName(v string) *CreateApiKeyResponseBodyDataCustomKeyList {
	s.KeyName = &v
	return s
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) SetLimitRate(v float32) *CreateApiKeyResponseBodyDataCustomKeyList {
	s.LimitRate = &v
	return s
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) SetLimitType(v string) *CreateApiKeyResponseBodyDataCustomKeyList {
	s.LimitType = &v
	return s
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) SetTokenQuota(v int64) *CreateApiKeyResponseBodyDataCustomKeyList {
	s.TokenQuota = &v
	return s
}

func (s *CreateApiKeyResponseBodyDataCustomKeyList) Validate() error {
	return dara.Validate(s)
}
