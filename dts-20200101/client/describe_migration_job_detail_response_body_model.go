// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) *DescribeMigrationJobDetailResponseBody
	GetDataInitializationDetailList() *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList
	SetDataSynchronizationDetailList(v *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) *DescribeMigrationJobDetailResponseBody
	GetDataSynchronizationDetailList() *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList
	SetErrCode(v string) *DescribeMigrationJobDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMigrationJobDetailResponseBody
	GetErrMessage() *string
	SetPageNumber(v int32) *DescribeMigrationJobDetailResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeMigrationJobDetailResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeMigrationJobDetailResponseBody
	GetRequestId() *string
	SetStructureInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) *DescribeMigrationJobDetailResponseBody
	GetStructureInitializationDetailList() *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList
	SetSuccess(v string) *DescribeMigrationJobDetailResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int64) *DescribeMigrationJobDetailResponseBody
	GetTotalRecordCount() *int64
}

type DescribeMigrationJobDetailResponseBody struct {
	DataInitializationDetailList  *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList  `json:"DataInitializationDetailList,omitempty" xml:"DataInitializationDetailList,omitempty" type:"Struct"`
	DataSynchronizationDetailList *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList `json:"DataSynchronizationDetailList,omitempty" xml:"DataSynchronizationDetailList,omitempty" type:"Struct"`
	// Specifies whether to query the details of incremental data migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > Default value: **false**
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Specifies whether to query the details of full data migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > Default value: **false**
	//
	// example:
	//
	// 0AE3CD0B-4148-426F-A90E-952467CC****
	RequestId                         *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StructureInitializationDetailList *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList `json:"StructureInitializationDetailList,omitempty" xml:"StructureInitializationDetailList,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBody) GetDataInitializationDetailList() *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList {
	return s.DataInitializationDetailList
}

func (s *DescribeMigrationJobDetailResponseBody) GetDataSynchronizationDetailList() *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList {
	return s.DataSynchronizationDetailList
}

func (s *DescribeMigrationJobDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMigrationJobDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMigrationJobDetailResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMigrationJobDetailResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeMigrationJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrationJobDetailResponseBody) GetStructureInitializationDetailList() *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList {
	return s.StructureInitializationDetailList
}

func (s *DescribeMigrationJobDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMigrationJobDetailResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataSynchronizationDetailList(v *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataSynchronizationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetErrCode(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetErrMessage(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageNumber(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetRequestId(v string) *DescribeMigrationJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetStructureInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.StructureInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetSuccess(v string) *DescribeMigrationJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobDetailResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) Validate() error {
	if s.DataInitializationDetailList != nil {
		if err := s.DataInitializationDetailList.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronizationDetailList != nil {
		if err := s.DataSynchronizationDetailList.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitializationDetailList != nil {
		if err := s.StructureInitializationDetailList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyDataInitializationDetailList struct {
	DataInitializationDetail []*DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail `json:"DataInitializationDetail,omitempty" xml:"DataInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) GetDataInitializationDetail() []*DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	return s.DataInitializationDetail
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) SetDataInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList {
	s.DataInitializationDetail = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) Validate() error {
	if s.DataInitializationDetail != nil {
		for _, item := range s.DataInitializationDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishRowNum           *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	MigrationTime          *string `json:"MigrationTime,omitempty" xml:"MigrationTime,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TotalRowNum            *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetFinishRowNum() *string {
	return s.FinishRowNum
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetMigrationTime() *string {
	return s.MigrationTime
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetTableName() *string {
	return s.TableName
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GetTotalRowNum() *string {
	return s.TotalRowNum
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetFinishRowNum(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.FinishRowNum = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetMigrationTime(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.MigrationTime = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetTableName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.TableName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetTotalRowNum(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.TotalRowNum = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList struct {
	DataSynchronizationDetail []*DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail `json:"DataSynchronizationDetail,omitempty" xml:"DataSynchronizationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) GetDataSynchronizationDetail() []*DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	return s.DataSynchronizationDetail
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) SetDataSynchronizationDetail(v []*DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList {
	s.DataSynchronizationDetail = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) Validate() error {
	if s.DataSynchronizationDetail != nil {
		for _, item := range s.DataSynchronizationDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GetTableName() *string {
	return s.TableName
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetTableName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.TableName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList struct {
	StructureInitializationDetail []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail `json:"StructureInitializationDetail,omitempty" xml:"StructureInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) GetStructureInitializationDetail() []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	return s.StructureInitializationDetail
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) SetStructureInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList {
	s.StructureInitializationDetail = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) Validate() error {
	if s.StructureInitializationDetail != nil {
		for _, item := range s.StructureInitializationDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail struct {
	ConstraintList         *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList `json:"ConstraintList,omitempty" xml:"ConstraintList,omitempty" type:"Struct"`
	DestinationOwnerDBName *string                                                                                                             `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string                                                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string                                                                                                             `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string                                                                                                             `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string                                                                                                             `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string                                                                                                             `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string                                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetConstraintList() *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList {
	return s.ConstraintList
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetObjectDefinition() *string {
	return s.ObjectDefinition
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetObjectName() *string {
	return s.ObjectName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetConstraintList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ConstraintList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) Validate() error {
	if s.ConstraintList != nil {
		if err := s.ConstraintList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList struct {
	StructureInitializationDetail []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail `json:"StructureInitializationDetail,omitempty" xml:"StructureInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) GetStructureInitializationDetail() []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	return s.StructureInitializationDetail
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) SetStructureInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList {
	s.StructureInitializationDetail = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) Validate() error {
	if s.StructureInitializationDetail != nil {
		for _, item := range s.StructureInitializationDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetObjectDefinition() *string {
	return s.ObjectDefinition
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetObjectName() *string {
	return s.ObjectName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) Validate() error {
	return dara.Validate(s)
}
