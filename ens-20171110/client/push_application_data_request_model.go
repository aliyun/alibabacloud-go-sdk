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
	// The application ID.
	//
	// Before calling this operation, activate the Edge Application Hosting service and create an application by calling CreateApplication. Obtain the AppId from the response of the CreateApplication operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of data files to be pre-pushed, in the string format of a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"app01\\",       \\"version\\":\\"1.0\\",       \\"size\\":100,\\"archiveType\\":\\"tar.gz\\",       \\"md5\\":\\"\\",       \\"url\\":\\"http://xxxx\\",\\"timeout\\": 1000   },    {       \\"name\\":\\"app02\\",       \\"version\\":\\"1.1\\",       \\"size\\":10,\\"archiveType\\":\\"zip\\",       \\"md5\\":\\"xxxx\\",       \\"url\\":\\"http://xxxxxx\\",\\"timeout\\": 1000   }]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The canary pre-push policy in JSON string format. Multiple distribution policies are supported. By default, all data is pre-pushed.
	//
	// example:
	//
	// {\\"name\\": \\"ScheduleToRegionId\\",  \\"parameters\\":{      \\"operator\\": \\"In\\",       \\"values\\": [\\"cn-chegndu-telecom-4\\", \\"cn-shanghai-cmcc-4\\"]  }}
	PushStrategy *string `json:"PushStrategy,omitempty" xml:"PushStrategy,omitempty"`
	// This parameter is invalid.
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
