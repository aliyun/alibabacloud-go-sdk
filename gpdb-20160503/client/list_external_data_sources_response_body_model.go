// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListExternalDataSourcesResponseBodyItems) *ListExternalDataSourcesResponseBody
	GetItems() []*ListExternalDataSourcesResponseBodyItems
	SetPageNumber(v int32) *ListExternalDataSourcesResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *ListExternalDataSourcesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListExternalDataSourcesResponseBody
	GetTotalRecordCount() *int32
}

type ListExternalDataSourcesResponseBody struct {
	// The Hadoop external table services.
	Items []*ListExternalDataSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// BBE00C04-A3E8-4114-881D-0480A72CB92E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListExternalDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalDataSourcesResponseBody) GetItems() []*ListExternalDataSourcesResponseBodyItems {
	return s.Items
}

func (s *ListExternalDataSourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExternalDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalDataSourcesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListExternalDataSourcesResponseBody) SetItems(v []*ListExternalDataSourcesResponseBodyItems) *ListExternalDataSourcesResponseBody {
	s.Items = v
	return s
}

func (s *ListExternalDataSourcesResponseBody) SetPageNumber(v int32) *ListExternalDataSourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExternalDataSourcesResponseBody) SetRequestId(v string) *ListExternalDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalDataSourcesResponseBody) SetTotalRecordCount(v int32) *ListExternalDataSourcesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListExternalDataSourcesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExternalDataSourcesResponseBodyItems struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The service directory in which Hadoop-related configuration files are stored.
	//
	// example:
	//
	// HadoopDir
	DataSourceDir *string `json:"DataSourceDir,omitempty" xml:"DataSourceDir,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 123
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// hdfs_pxf
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
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
	// Running
	DataSourceStatus *string `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// HDFS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The Id of external data service
	//
	// example:
	//
	// 123
	ExternalDataServiceId *int32 `json:"ExternalDataServiceId,omitempty" xml:"ExternalDataServiceId,omitempty"`
	// The time when the service was last modified.
	//
	// example:
	//
	// 2019-10-08T16:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The information about the service status. For example, if the service is in the exception state, the cause of the exception is displayed. If the service is in the running state, this parameter is left empty.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListExternalDataSourcesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataSourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListExternalDataSourcesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceDir() *string {
	return s.DataSourceDir
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *ListExternalDataSourcesResponseBodyItems) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListExternalDataSourcesResponseBodyItems) GetExternalDataServiceId() *int32 {
	return s.ExternalDataServiceId
}

func (s *ListExternalDataSourcesResponseBodyItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListExternalDataSourcesResponseBodyItems) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListExternalDataSourcesResponseBodyItems) SetCreateTime(v string) *ListExternalDataSourcesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceDescription(v string) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceDescription = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceDir(v string) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceDir = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceId(v int32) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceId = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceName(v string) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceName = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceStatus(v string) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceStatus = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetDataSourceType(v string) *ListExternalDataSourcesResponseBodyItems {
	s.DataSourceType = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetExternalDataServiceId(v int32) *ListExternalDataSourcesResponseBodyItems {
	s.ExternalDataServiceId = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetModifyTime(v string) *ListExternalDataSourcesResponseBodyItems {
	s.ModifyTime = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) SetStatusMessage(v string) *ListExternalDataSourcesResponseBodyItems {
	s.StatusMessage = &v
	return s
}

func (s *ListExternalDataSourcesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
