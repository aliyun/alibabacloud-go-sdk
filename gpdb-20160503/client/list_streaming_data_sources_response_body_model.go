// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceItems(v []*ListStreamingDataSourcesResponseBodyDataSourceItems) *ListStreamingDataSourcesResponseBody
	GetDataSourceItems() []*ListStreamingDataSourcesResponseBodyDataSourceItems
	SetPageNumber(v int32) *ListStreamingDataSourcesResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *ListStreamingDataSourcesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListStreamingDataSourcesResponseBody
	GetTotalRecordCount() *int32
}

type ListStreamingDataSourcesResponseBody struct {
	// The queried data sources.
	DataSourceItems []*ListStreamingDataSourcesResponseBodyDataSourceItems `json:"DataSourceItems,omitempty" xml:"DataSourceItems,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListStreamingDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamingDataSourcesResponseBody) GetDataSourceItems() []*ListStreamingDataSourcesResponseBodyDataSourceItems {
	return s.DataSourceItems
}

func (s *ListStreamingDataSourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStreamingDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamingDataSourcesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListStreamingDataSourcesResponseBody) SetDataSourceItems(v []*ListStreamingDataSourcesResponseBodyDataSourceItems) *ListStreamingDataSourcesResponseBody {
	s.DataSourceItems = v
	return s
}

func (s *ListStreamingDataSourcesResponseBody) SetPageNumber(v int32) *ListStreamingDataSourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBody) SetRequestId(v string) *ListStreamingDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBody) SetTotalRecordCount(v int32) *ListStreamingDataSourcesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStreamingDataSourcesResponseBodyDataSourceItems struct {
	// The time when the data source was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configurations of the data source.
	//
	// example:
	//
	// {"brokers":"broker0:9091,broker1:9091","topic":"topic"}
	DataSourceConfig *string `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test-kafka
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 1
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test-kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- kafka
	//
	// example:
	//
	// kafka
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The information about the service status. For example, if the service is in the exception state, the cause of the exception is displayed. If the service is in the running state, this parameter is left empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the data source was last modified.
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 1
	ServiceId *int32 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- init
	//
	// 	- running
	//
	// 	- exception
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListStreamingDataSourcesResponseBodyDataSourceItems) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataSourcesResponseBodyDataSourceItems) GoString() string {
	return s.String()
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetServiceId() *int32 {
	return s.ServiceId
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) GetStatus() *string {
	return s.Status
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetCreateTime(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.CreateTime = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetDataSourceConfig(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.DataSourceConfig = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetDataSourceDescription(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.DataSourceDescription = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetDataSourceId(v int32) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.DataSourceId = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetDataSourceName(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.DataSourceName = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetDataSourceType(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.DataSourceType = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetErrorMessage(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetModifyTime(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.ModifyTime = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetServiceId(v int32) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.ServiceId = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) SetStatus(v string) *ListStreamingDataSourcesResponseBodyDataSourceItems {
	s.Status = &v
	return s
}

func (s *ListStreamingDataSourcesResponseBodyDataSourceItems) Validate() error {
	return dara.Validate(s)
}
