// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageCloudMonitorConfig(v *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) *GetMessageCloudMonitorConfigResponseBody
	GetMessageCloudMonitorConfig() *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig
	SetRequestId(v string) *GetMessageCloudMonitorConfigResponseBody
	GetRequestId() *string
}

type GetMessageCloudMonitorConfigResponseBody struct {
	MessageCloudMonitorConfig *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig `json:"MessageCloudMonitorConfig,omitempty" xml:"MessageCloudMonitorConfig,omitempty" type:"Struct"`
	RequestId                 *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageCloudMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorConfigResponseBody) GetMessageCloudMonitorConfig() *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig {
	return s.MessageCloudMonitorConfig
}

func (s *GetMessageCloudMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageCloudMonitorConfigResponseBody) SetMessageCloudMonitorConfig(v *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) *GetMessageCloudMonitorConfigResponseBody {
	s.MessageCloudMonitorConfig = v
	return s
}

func (s *GetMessageCloudMonitorConfigResponseBody) SetRequestId(v string) *GetMessageCloudMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageCloudMonitorConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
}

func (s GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) SetAppId(v string) *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig {
	s.AppId = &v
	return s
}

func (s *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) SetEventTypeList(v string) *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig {
	s.EventTypeList = &v
	return s
}

func (s *GetMessageCloudMonitorConfigResponseBodyMessageCloudMonitorConfig) Validate() error {
	return dara.Validate(s)
}
