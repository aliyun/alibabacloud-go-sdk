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
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeRestoreRangeInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E2BD9DFC-6760-5F49-97C5-DA739E29****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
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
	// example:
	//
	// 127.0.0.1
	BackupSourceHost *string `json:"BackupSourceHost,omitempty" xml:"BackupSourceHost,omitempty"`
	// example:
	//
	// rm-testxx
	BackupSourceInstanceId *string `json:"BackupSourceInstanceId,omitempty" xml:"BackupSourceInstanceId,omitempty"`
	// example:
	//
	// rds
	BackupSourceInstanceType *string `json:"BackupSourceInstanceType,omitempty" xml:"BackupSourceInstanceType,omitempty"`
	// example:
	//
	// 3306
	BackupSourcePort           *string                                                                 `json:"BackupSourcePort,omitempty" xml:"BackupSourcePort,omitempty"`
	BeginTimestampForRestore   *int64                                                                  `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	EndTimestampForRestore     *int64                                                                  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	FullBackupList             *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList `json:"FullBackupList,omitempty" xml:"FullBackupList,omitempty" type:"Struct"`
	RangeType                  *string                                                                 `json:"RangeType,omitempty" xml:"RangeType,omitempty"`
	SourceEndpointInstanceID   *string                                                                 `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointInstanceType *string                                                                 `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetBackupSourceHost() *string {
	return s.BackupSourceHost
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetBackupSourceInstanceId() *string {
	return s.BackupSourceInstanceId
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetBackupSourceInstanceType() *string {
	return s.BackupSourceInstanceType
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GetBackupSourcePort() *string {
	return s.BackupSourcePort
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

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBackupSourceHost(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BackupSourceHost = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBackupSourceInstanceId(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BackupSourceInstanceId = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBackupSourceInstanceType(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BackupSourceInstanceType = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBackupSourcePort(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BackupSourcePort = &v
	return s
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
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
