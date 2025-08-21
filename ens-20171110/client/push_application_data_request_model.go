// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushApplicationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PushApplicationDataRequest
	GetAppId() *string
	SetData(v string) *PushApplicationDataRequest
	GetData() *string
	SetPushStrategy(v string) *PushApplicationDataRequest
	GetPushStrategy() *string
	SetTimeout(v int32) *PushApplicationDataRequest
	GetTimeout() *int32
}

type PushApplicationDataRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data files that you want to push. The value must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"app01\\",       \\"version\\":\\"1.0\\",       \\"size\\":100,\\"archiveType\\":\\"tar.gz\\",       \\"md5\\":\\"\\",       \\"url\\":\\"http://xxxx\\",\\"timeout\\": 1000   },    {       \\"name\\":\\"app02\\",       \\"version\\":\\"1.1\\",       \\"size\\":10,\\"archiveType\\":\\"zip\\",       \\"md5\\":\\"xxxx\\",       \\"url\\":\\"http://xxxxxx\\",\\"timeout\\": 1000   }]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The push policy in the canary release environment. The value must be a JSON string. You can specify multiple push policies. By default, all data files are pushed.
	//
	// example:
	//
	// {\\"name\\": \\"ScheduleToRegionId\\",  \\"parameters\\":{      \\"operator\\": \\"In\\",       \\"values\\": [\\"cn-chegndu-telecom-4\\", \\"cn-shanghai-cmcc-4\\"]  }}
	PushStrategy *string `json:"PushStrategy,omitempty" xml:"PushStrategy,omitempty"`
	// This parameter does not take effect.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s PushApplicationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PushApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *PushApplicationDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushApplicationDataRequest) GetData() *string {
	return s.Data
}

func (s *PushApplicationDataRequest) GetPushStrategy() *string {
	return s.PushStrategy
}

func (s *PushApplicationDataRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *PushApplicationDataRequest) SetAppId(v string) *PushApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *PushApplicationDataRequest) SetData(v string) *PushApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *PushApplicationDataRequest) SetPushStrategy(v string) *PushApplicationDataRequest {
	s.PushStrategy = &v
	return s
}

func (s *PushApplicationDataRequest) SetTimeout(v int32) *PushApplicationDataRequest {
	s.Timeout = &v
	return s
}

func (s *PushApplicationDataRequest) Validate() error {
	return dara.Validate(s)
}
