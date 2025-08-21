// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistApplicationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DistApplicationDataRequest
	GetAppId() *string
	SetData(v string) *DistApplicationDataRequest
	GetData() *string
	SetDistStrategy(v string) *DistApplicationDataRequest
	GetDistStrategy() *string
}

type DistApplicationDataRequest struct {
	// The ID of the application. To obtain the application ID, call the ListApplications operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of data files that you want to distribute. The value must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"app01\\",        \\"version\\":\\"1.0\\",        \\"destPath\\":\\"/root/installed\\",        \\"decompress\\":true,        \\"targetDirName\\":\\"target01\\",        \\"fileMode\\":755,        \\"timeout\\":1000    },    {        \\"name\\":\\"app02\\",        \\"version\\":\\"1.1\\",        \\"destPath\\":\\"/tmp/test.txt\\",        \\"decompress\\":false    }]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The canary release policy. The value must be a JSON string. You can specify multiple distribution policies. By default, all data files are distributed.
	//
	// example:
	//
	// {\\"name\\":\\"ScheduleToAllByMatchExpressions\\",\\"parameters\\":{\\"match_expressions\\":[{\\"key\\":\\"region_id\\",\\"operator\\":\\"In\\",\\"values\\":[\\"cn-wuhan-telecom_unicom_cmcc-2\\"]}]}}
	DistStrategy *string `json:"DistStrategy,omitempty" xml:"DistStrategy,omitempty"`
}

func (s DistApplicationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *DistApplicationDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DistApplicationDataRequest) GetData() *string {
	return s.Data
}

func (s *DistApplicationDataRequest) GetDistStrategy() *string {
	return s.DistStrategy
}

func (s *DistApplicationDataRequest) SetAppId(v string) *DistApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *DistApplicationDataRequest) SetData(v string) *DistApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *DistApplicationDataRequest) SetDistStrategy(v string) *DistApplicationDataRequest {
	s.DistStrategy = &v
	return s
}

func (s *DistApplicationDataRequest) Validate() error {
	return dara.Validate(s)
}
