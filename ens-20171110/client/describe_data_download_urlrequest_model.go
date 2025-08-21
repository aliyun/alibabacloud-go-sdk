// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDownloadURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeDataDownloadURLRequest
	GetAppId() *string
	SetDataName(v string) *DescribeDataDownloadURLRequest
	GetDataName() *string
	SetDataVersion(v string) *DescribeDataDownloadURLRequest
	GetDataVersion() *string
	SetExpireTimeout(v int64) *DescribeDataDownloadURLRequest
	GetExpireTimeout() *int64
	SetServerFilterStrategy(v string) *DescribeDataDownloadURLRequest
	GetServerFilterStrategy() *string
}

type DescribeDataDownloadURLRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 474bdef0-d149-4695-abfb-52912d9143f0
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the data file.
	//
	// This parameter is required.
	//
	// example:
	//
	// mirror_file/pk-1642597182026-878199448832413.tar
	DataName *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	// The version number of the data file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7895
	DataVersion *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// 600
	ExpireTimeout *int64 `json:"ExpireTimeout,omitempty" xml:"ExpireTimeout,omitempty"`
	// The condition that you want to use to filter file servers. You can specify multiple canary release policies. By default, all resources are queried.
	//
	// example:
	//
	// {\\"name\\": \\"ScheduleToRegionId\\",\\"parameters\\":{\\"operator\\": \\"In\\",\\"values\\": [\\"cn-shijiazhuang-telecom_unicom_cmcc\\"]}}
	ServerFilterStrategy *string `json:"ServerFilterStrategy,omitempty" xml:"ServerFilterStrategy,omitempty"`
}

func (s DescribeDataDownloadURLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDownloadURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDataDownloadURLRequest) GetDataName() *string {
	return s.DataName
}

func (s *DescribeDataDownloadURLRequest) GetDataVersion() *string {
	return s.DataVersion
}

func (s *DescribeDataDownloadURLRequest) GetExpireTimeout() *int64 {
	return s.ExpireTimeout
}

func (s *DescribeDataDownloadURLRequest) GetServerFilterStrategy() *string {
	return s.ServerFilterStrategy
}

func (s *DescribeDataDownloadURLRequest) SetAppId(v string) *DescribeDataDownloadURLRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetDataName(v string) *DescribeDataDownloadURLRequest {
	s.DataName = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetDataVersion(v string) *DescribeDataDownloadURLRequest {
	s.DataVersion = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetExpireTimeout(v int64) *DescribeDataDownloadURLRequest {
	s.ExpireTimeout = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetServerFilterStrategy(v string) *DescribeDataDownloadURLRequest {
	s.ServerFilterStrategy = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) Validate() error {
	return dara.Validate(s)
}
