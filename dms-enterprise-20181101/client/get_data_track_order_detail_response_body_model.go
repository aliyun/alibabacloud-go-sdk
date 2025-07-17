// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataTrackOrderDetail(v *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) *GetDataTrackOrderDetailResponseBody
	GetDataTrackOrderDetail() *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail
	SetErrorCode(v string) *GetDataTrackOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataTrackOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataTrackOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataTrackOrderDetailResponseBody
	GetSuccess() *bool
}

type GetDataTrackOrderDetailResponseBody struct {
	// The details of the ticket.
	DataTrackOrderDetail *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail `json:"DataTrackOrderDetail,omitempty" xml:"DataTrackOrderDetail,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataTrackOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataTrackOrderDetailResponseBody) GetDataTrackOrderDetail() *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	return s.DataTrackOrderDetail
}

func (s *GetDataTrackOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataTrackOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataTrackOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataTrackOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataTrackOrderDetailResponseBody) SetDataTrackOrderDetail(v *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) *GetDataTrackOrderDetailResponseBody {
	s.DataTrackOrderDetail = v
	return s
}

func (s *GetDataTrackOrderDetailResponseBody) SetErrorCode(v string) *GetDataTrackOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBody) SetErrorMessage(v string) *GetDataTrackOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBody) SetRequestId(v string) *GetDataTrackOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBody) SetSuccess(v bool) *GetDataTrackOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail struct {
	// The name that is used to search for the database.
	//
	// example:
	//
	// xxx@yyy:3306
	DatabaseSearchName *string `json:"DatabaseSearchName,omitempty" xml:"DatabaseSearchName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 3431****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The end time of the time range in which data operations are tracked. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 10:00:00
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// The start time of the time range in which data operations are tracked. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 00:00:00
	JobStartTime *string `json:"JobStartTime,omitempty" xml:"JobStartTime,omitempty"`
	// The status of the data tracking task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **LISTING**: The binary logs are being obtained.
	//
	// 	- **LIST_SUCCESS**: The binary logs are successfully obtained.
	//
	// 	- **DOWNLOADING**: The binary logs are being downloaded.
	//
	// 	- **DOWNLOAD_FAIL**: The binary logs failed to be downloaded.
	//
	// 	- **DOWNLOAD_SUCCESS**: The binary logs are successfully downloaded.
	//
	// 	- **FILTERING**: The binary logs are being parsed.
	//
	// 	- **FILTER_FAIL**: The binary logs failed to be parsed.
	//
	// 	- **FILTER_SUCCESS**: The binary logs are successfully parsed.
	//
	// example:
	//
	// FILTER_SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// as_task
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The description of the task status.
	//
	// example:
	//
	// searching success
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The names of the tables for which data operations are tracked.
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
	// The types of data operations that are tracked.
	TrackTypes []*string `json:"TrackTypes,omitempty" xml:"TrackTypes,omitempty" type:"Repeated"`
}

func (s GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetDbId() *int64 {
	return s.DbId
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetJobEndTime() *string {
	return s.JobEndTime
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetJobStartTime() *string {
	return s.JobStartTime
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetLogic() *bool {
	return s.Logic
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetTableNames() []*string {
	return s.TableNames
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) GetTrackTypes() []*string {
	return s.TrackTypes
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetDatabaseSearchName(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.DatabaseSearchName = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetDbId(v int64) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.DbId = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetJobEndTime(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.JobEndTime = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetJobStartTime(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.JobStartTime = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetJobStatus(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.JobStatus = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetLogic(v bool) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.Logic = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetSchemaName(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.SchemaName = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetStatusDesc(v string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.StatusDesc = &v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetTableNames(v []*string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.TableNames = v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) SetTrackTypes(v []*string) *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail {
	s.TrackTypes = v
	return s
}

func (s *GetDataTrackOrderDetailResponseBodyDataTrackOrderDetail) Validate() error {
	return dara.Validate(s)
}
