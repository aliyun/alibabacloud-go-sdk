// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *GetStorageNotifyConfigResponseBody
	GetConfigType() *string
	SetOssNotifyConfig(v *GetStorageNotifyConfigResponseBodyOssNotifyConfig) *GetStorageNotifyConfigResponseBody
	GetOssNotifyConfig() *GetStorageNotifyConfigResponseBodyOssNotifyConfig
	SetRequestId(v string) *GetStorageNotifyConfigResponseBody
	GetRequestId() *string
}

type GetStorageNotifyConfigResponseBody struct {
	ConfigType      *string                                            `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	OssNotifyConfig *GetStorageNotifyConfigResponseBodyOssNotifyConfig `json:"OssNotifyConfig,omitempty" xml:"OssNotifyConfig,omitempty" type:"Struct"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetStorageNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageNotifyConfigResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetStorageNotifyConfigResponseBody) GetOssNotifyConfig() *GetStorageNotifyConfigResponseBodyOssNotifyConfig {
	return s.OssNotifyConfig
}

func (s *GetStorageNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageNotifyConfigResponseBody) SetConfigType(v string) *GetStorageNotifyConfigResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBody) SetOssNotifyConfig(v *GetStorageNotifyConfigResponseBodyOssNotifyConfig) *GetStorageNotifyConfigResponseBody {
	s.OssNotifyConfig = v
	return s
}

func (s *GetStorageNotifyConfigResponseBody) SetRequestId(v string) *GetStorageNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageNotifyConfigResponseBodyOssNotifyConfig struct {
	EventList          *string `json:"EventList,omitempty" xml:"EventList,omitempty"`
	HttpProcessAddress *string `json:"HttpProcessAddress,omitempty" xml:"HttpProcessAddress,omitempty"`
	NotifyName         *string `json:"NotifyName,omitempty" xml:"NotifyName,omitempty"`
	ResourcePrefixName *string `json:"ResourcePrefixName,omitempty" xml:"ResourcePrefixName,omitempty"`
}

func (s GetStorageNotifyConfigResponseBodyOssNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStorageNotifyConfigResponseBodyOssNotifyConfig) GoString() string {
	return s.String()
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) GetEventList() *string {
	return s.EventList
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) GetHttpProcessAddress() *string {
	return s.HttpProcessAddress
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) GetNotifyName() *string {
	return s.NotifyName
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) GetResourcePrefixName() *string {
	return s.ResourcePrefixName
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) SetEventList(v string) *GetStorageNotifyConfigResponseBodyOssNotifyConfig {
	s.EventList = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) SetHttpProcessAddress(v string) *GetStorageNotifyConfigResponseBodyOssNotifyConfig {
	s.HttpProcessAddress = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) SetNotifyName(v string) *GetStorageNotifyConfigResponseBodyOssNotifyConfig {
	s.NotifyName = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) SetResourcePrefixName(v string) *GetStorageNotifyConfigResponseBodyOssNotifyConfig {
	s.ResourcePrefixName = &v
	return s
}

func (s *GetStorageNotifyConfigResponseBodyOssNotifyConfig) Validate() error {
	return dara.Validate(s)
}
