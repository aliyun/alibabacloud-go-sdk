// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeStreamingDataSourceResponseBody
	GetCreateTime() *string
	SetDataSourceConfig(v string) *DescribeStreamingDataSourceResponseBody
	GetDataSourceConfig() *string
	SetDataSourceDescription(v string) *DescribeStreamingDataSourceResponseBody
	GetDataSourceDescription() *string
	SetDataSourceId(v string) *DescribeStreamingDataSourceResponseBody
	GetDataSourceId() *string
	SetDataSourceName(v string) *DescribeStreamingDataSourceResponseBody
	GetDataSourceName() *string
	SetDataSourceType(v string) *DescribeStreamingDataSourceResponseBody
	GetDataSourceType() *string
	SetErrorMessage(v string) *DescribeStreamingDataSourceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeStreamingDataSourceResponseBody
	GetRequestId() *string
	SetServiceId(v int32) *DescribeStreamingDataSourceResponseBody
	GetServiceId() *int32
	SetStatus(v string) *DescribeStreamingDataSourceResponseBody
	GetStatus() *string
}

type DescribeStreamingDataSourceResponseBody struct {
	// Creation time.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Data source configuration information.
	//
	// example:
	//
	// {"brokers":"broker0:9091,broker1:9091","topic":"topic"}
	DataSourceConfig *string `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty"`
	// Data source description.
	//
	// example:
	//
	// test-kafka
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// 1
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Data source name.
	//
	// example:
	//
	// test-kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Data source type, values include:
	//
	//  -  kafka
	//
	// example:
	//
	// kafka
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// Service status message, for example, in case of an exception, it will show the reason for the exception. In normal Running state, this value is empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// External data service ID.
	//
	// example:
	//
	// 1
	ServiceId *int32 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service status:
	//
	// - Initializing init
	//
	// - Running running
	//
	// - Exception exception
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeStreamingDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataSourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeStreamingDataSourceResponseBody) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *DescribeStreamingDataSourceResponseBody) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *DescribeStreamingDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeStreamingDataSourceResponseBody) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeStreamingDataSourceResponseBody) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeStreamingDataSourceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeStreamingDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamingDataSourceResponseBody) GetServiceId() *int32 {
	return s.ServiceId
}

func (s *DescribeStreamingDataSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeStreamingDataSourceResponseBody) SetCreateTime(v string) *DescribeStreamingDataSourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetDataSourceConfig(v string) *DescribeStreamingDataSourceResponseBody {
	s.DataSourceConfig = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetDataSourceDescription(v string) *DescribeStreamingDataSourceResponseBody {
	s.DataSourceDescription = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetDataSourceId(v string) *DescribeStreamingDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetDataSourceName(v string) *DescribeStreamingDataSourceResponseBody {
	s.DataSourceName = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetDataSourceType(v string) *DescribeStreamingDataSourceResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetErrorMessage(v string) *DescribeStreamingDataSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetRequestId(v string) *DescribeStreamingDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetServiceId(v int32) *DescribeStreamingDataSourceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) SetStatus(v string) *DescribeStreamingDataSourceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamingDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
