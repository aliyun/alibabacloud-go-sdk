// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreRangeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeRestoreRangeInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeRestoreRangeInfoResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeRestoreRangeInfoResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeRestoreRangeInfoResponseBodyItems) *DescribeRestoreRangeInfoResponseBody
	GetItems() *DescribeRestoreRangeInfoResponseBodyItems
	SetRequestId(v string) *DescribeRestoreRangeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRestoreRangeInfoResponseBody
	GetSuccess() *bool
}

type DescribeRestoreRangeInfoResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The information about the time ranges to which you can restore data.
	Items *DescribeRestoreRangeInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E2BD9DFC-6760-5F49-97C5-DA739E29****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeRestoreRangeInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeRestoreRangeInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeRestoreRangeInfoResponseBody) GetItems() *DescribeRestoreRangeInfoResponseBodyItems {
	return s.Items
}

func (s *DescribeRestoreRangeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreRangeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRestoreRangeInfoResponseBody) SetErrCode(v string) *DescribeRestoreRangeInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetErrMessage(v string) *DescribeRestoreRangeInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetHttpStatusCode(v int32) *DescribeRestoreRangeInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetItems(v *DescribeRestoreRangeInfoResponseBodyItems) *DescribeRestoreRangeInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetRequestId(v string) *DescribeRestoreRangeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetSuccess(v bool) *DescribeRestoreRangeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreRangeInfoResponseBodyItems struct {
	DBSRecoverRange []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange `json:"DBSRecoverRange,omitempty" xml:"DBSRecoverRange,omitempty" type:"Repeated"`
}

func (s DescribeRestoreRangeInfoResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItems) GetDBSRecoverRange() []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	return s.DBSRecoverRange
}

func (s *DescribeRestoreRangeInfoResponseBodyItems) SetDBSRecoverRange(v []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) *DescribeRestoreRangeInfoResponseBodyItems {
	s.DBSRecoverRange = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItems) Validate() error {
	if s.DBSRecoverRange != nil {
		for _, item := range s.DBSRecoverRange {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange struct {
	// The beginning of the time range to which you can restore data.
	//
	// example:
	//
	// 1646760282000
	BeginTimestampForRestore *int64 `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	// The end of the time range to which you can restore data.
	//
	// example:
	//
	// 1646760308000
	EndTimestampForRestore *int64 `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	// If the value of the RangeType parameter is point, this parameter is returned. The value of this parameter describes information about all backup points in the time range.
	FullBackupList *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList `json:"FullBackupList,omitempty" xml:"FullBackupList,omitempty" type:"Struct"`
	// The type of the time range to which you can restore data.
	//
	// 	- **point**: The time range contains discrete points in time at which full backups were performed.
	//
	// 	- **range**: The time range is a period of time for which continuous backup is performed. You can specify a random point in time in the time range to restore data.
	//
	// example:
	//
	// point
	RangeType *string `json:"RangeType,omitempty" xml:"RangeType,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// rm-bp106x9tk2c91****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database.
	//
	// example:
	//
	// rds
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetBeginTimestampForRestore() *int64 {
	return s.BeginTimestampForRestore
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetEndTimestampForRestore() *int64 {
	return s.EndTimestampForRestore
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetFullBackupList() *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList {
	return s.FullBackupList
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetRangeType() *string {
	return s.RangeType
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBeginTimestampForRestore(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetEndTimestampForRestore(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetFullBackupList(v *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.FullBackupList = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetRangeType(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.RangeType = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetSourceEndpointInstanceID(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetSourceEndpointInstanceType(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) Validate() error {
	if s.FullBackupList != nil {
		if err := s.FullBackupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList struct {
	FullBackupDetail []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail `json:"FullBackupDetail,omitempty" xml:"FullBackupDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) GetFullBackupDetail() []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	return s.FullBackupDetail
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) SetFullBackupDetail(v []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList {
	s.FullBackupDetail = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) Validate() error {
	if s.FullBackupDetail != nil {
		for _, item := range s.FullBackupDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail struct {
	// The ID of the backup set.
	//
	// example:
	//
	// qecnsxkd****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The end time of the full backup task. Example: 1646760308000.
	//
	// example:
	//
	// 1646760308000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the full backup task. Example: 1646760282000.
	//
	// example:
	//
	// 1646760282000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetBackupSetId(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.BackupSetId = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetEndTime(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetStartTime(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) Validate() error {
	return dara.Validate(s)
}
