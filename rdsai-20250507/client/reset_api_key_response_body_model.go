// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResetApiKeyResponseBodyData) *ResetApiKeyResponseBody
	GetData() *ResetApiKeyResponseBodyData
	SetMessage(v string) *ResetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetApiKeyResponseBody
	GetSuccess() *bool
}

type ResetApiKeyResponseBody struct {
	Data *ResetApiKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBody) GetData() *ResetApiKeyResponseBodyData {
	return s.Data
}

func (s *ResetApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetApiKeyResponseBody) SetData(v *ResetApiKeyResponseBodyData) *ResetApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ResetApiKeyResponseBody) SetMessage(v string) *ResetApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetRequestId(v string) *ResetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetSuccess(v bool) *ResetApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ResetApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetApiKeyResponseBodyData struct {
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// http://xxx.yy/v1
	BaseUrl       *string                                     `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	CustomKeyList []*ResetApiKeyResponseBodyDataCustomKeyList `json:"CustomKeyList,omitempty" xml:"CustomKeyList,omitempty" type:"Repeated"`
}

func (s ResetApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBodyData) GetApiKey() *string {
	return s.ApiKey
}

func (s *ResetApiKeyResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ResetApiKeyResponseBodyData) GetCustomKeyList() []*ResetApiKeyResponseBodyDataCustomKeyList {
	return s.CustomKeyList
}

func (s *ResetApiKeyResponseBodyData) SetApiKey(v string) *ResetApiKeyResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *ResetApiKeyResponseBodyData) SetBaseUrl(v string) *ResetApiKeyResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *ResetApiKeyResponseBodyData) SetCustomKeyList(v []*ResetApiKeyResponseBodyDataCustomKeyList) *ResetApiKeyResponseBodyData {
	s.CustomKeyList = v
	return s
}

func (s *ResetApiKeyResponseBodyData) Validate() error {
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

type ResetApiKeyResponseBodyDataCustomKeyList struct {
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s ResetApiKeyResponseBodyDataCustomKeyList) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBodyDataCustomKeyList) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBodyDataCustomKeyList) GetApiKey() *string {
	return s.ApiKey
}

func (s *ResetApiKeyResponseBodyDataCustomKeyList) GetKeyName() *string {
	return s.KeyName
}

func (s *ResetApiKeyResponseBodyDataCustomKeyList) SetApiKey(v string) *ResetApiKeyResponseBodyDataCustomKeyList {
	s.ApiKey = &v
	return s
}

func (s *ResetApiKeyResponseBodyDataCustomKeyList) SetKeyName(v string) *ResetApiKeyResponseBodyDataCustomKeyList {
	s.KeyName = &v
	return s
}

func (s *ResetApiKeyResponseBodyDataCustomKeyList) Validate() error {
	return dara.Validate(s)
}
