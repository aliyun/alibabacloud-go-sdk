// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitializationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInitializationDetails(v []*DescribeInitializationStatusResponseBodyDataInitializationDetails) *DescribeInitializationStatusResponseBody
	GetDataInitializationDetails() []*DescribeInitializationStatusResponseBodyDataInitializationDetails
	SetDataSynchronizationDetails(v []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails) *DescribeInitializationStatusResponseBody
	GetDataSynchronizationDetails() []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails
	SetErrCode(v string) *DescribeInitializationStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInitializationStatusResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DescribeInitializationStatusResponseBody
	GetRequestId() *string
	SetStructureInitializationDetails(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetails) *DescribeInitializationStatusResponseBody
	GetStructureInitializationDetails() []*DescribeInitializationStatusResponseBodyStructureInitializationDetails
	SetSuccess(v string) *DescribeInitializationStatusResponseBody
	GetSuccess() *string
}

type DescribeInitializationStatusResponseBody struct {
	// The details of initial full data synchronization.
	DataInitializationDetails []*DescribeInitializationStatusResponseBodyDataInitializationDetails `json:"DataInitializationDetails,omitempty" xml:"DataInitializationDetails,omitempty" type:"Repeated"`
	// The details of incremental data synchronization.
	//
	// >  This parameter and the parameters it contains will be removed in the future.
	DataSynchronizationDetails []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails `json:"DataSynchronizationDetails,omitempty" xml:"DataSynchronizationDetails,omitempty" type:"Repeated"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 150DECD9-13FF-4929-A5DE-855BE9CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of initial schema synchronization.
	StructureInitializationDetails []*DescribeInitializationStatusResponseBodyStructureInitializationDetails `json:"StructureInitializationDetails,omitempty" xml:"StructureInitializationDetails,omitempty" type:"Repeated"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInitializationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBody) GetDataInitializationDetails() []*DescribeInitializationStatusResponseBodyDataInitializationDetails {
	return s.DataInitializationDetails
}

func (s *DescribeInitializationStatusResponseBody) GetDataSynchronizationDetails() []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	return s.DataSynchronizationDetails
}

func (s *DescribeInitializationStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInitializationStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInitializationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInitializationStatusResponseBody) GetStructureInitializationDetails() []*DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	return s.StructureInitializationDetails
}

func (s *DescribeInitializationStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeInitializationStatusResponseBody) SetDataInitializationDetails(v []*DescribeInitializationStatusResponseBodyDataInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.DataInitializationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetDataSynchronizationDetails(v []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails) *DescribeInitializationStatusResponseBody {
	s.DataSynchronizationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrCode(v string) *DescribeInitializationStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrMessage(v string) *DescribeInitializationStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetRequestId(v string) *DescribeInitializationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetStructureInitializationDetails(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.StructureInitializationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetSuccess(v string) *DescribeInitializationStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) Validate() error {
	if s.DataInitializationDetails != nil {
		for _, item := range s.DataInitializationDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataSynchronizationDetails != nil {
		for _, item := range s.DataSynchronizationDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StructureInitializationDetails != nil {
		for _, item := range s.StructureInitializationDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInitializationStatusResponseBodyDataInitializationDetails struct {
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	// The error message returned if initial full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The total number of rows that are actually synchronized.
	//
	// >  This parameter indicates the total number of actually synchronized rows. In contrast, the value of the **TotalRowNum*	- parameter is calculated based on the system tables in the source database. The values of the two parameters may be different due to time difference.
	//
	// example:
	//
	// 9993
	FinishRowNum *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	SourceOwnerDBName *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	// The status of initial full data synchronization. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Migrating**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The table name.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of rows that are supposed to be synchronized.
	//
	// >  The value of this parameter is calculated based on the system tables in the source database. In contrast, the **FinishRowNum*	- parameter indicates the total number of actually synchronized rows. The values of the two parameters may be different due to time difference.
	//
	// example:
	//
	// 9981
	TotalRowNum *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
	// The time spent on full data synchronization.
	//
	// example:
	//
	// 0.0
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetFinishRowNum() *string {
	return s.FinishRowNum
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetTableName() *string {
	return s.TableName
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetTotalRowNum() *string {
	return s.TotalRowNum
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) GetUsedTime() *string {
	return s.UsedTime
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetFinishRowNum(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.FinishRowNum = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TableName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTotalRowNum(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TotalRowNum = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetUsedTime(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.UsedTime = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeInitializationStatusResponseBodyDataSynchronizationDetails struct {
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	// The error message returned if incremental data synchronization failed.
	//
	// example:
	//
	// The task has failed for too long and cannot be repaired
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	SourceOwnerDBName *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	// The status of incremental data synchronization. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Migrating**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The table name.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GetTableName() *string {
	return s.TableName
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.TableName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetails struct {
	// The constraints of the synchronization object, such as indexes and foreign keys.
	//
	// >  This parameter is returned only if the **ObjectType*	- parameter is set to **Table*	- and the synchronization object has constraints.
	Constraints []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Repeated"`
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	// The error message returned if initial schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: Table \\"customer\\" already exists
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The schema of the object.
	//
	// example:
	//
	// CREATE TABLE `dtstestdata`.`customer` (\\n`customer_id`  int(10) unsigned   auto_increment  COMMENT \\"\\"   NOT NULL   , \\n`customer_title`  varchar(100)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NOT NULL   , \\n`customer_company1216`  varchar(40)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NOT NULL   , \\n`submission_date1216`  date     COMMENT \\"\\"   NULL   \\n, PRIMARY KEY (`customer_id`)) engine=InnoDB AUTO_INCREMENT=200001 DEFAULT CHARSET=`utf8` DEFAULT COLLATE `utf8_general_ci` ROW_FORMAT= Dynamic comment = \\"\\" ;\\n
	ObjectDefinition *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// customer
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The type of the object. Valid values:
	//
	// **Table**, **Constraint**, **Index**, **View**, **Materialize View**, **Type**, **Synonym**, **Trigger**, **Function**, **Procedure**, **Package**, **Default**, **Rule**, **PlanGuide**, and **Sequence**.
	//
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	SourceOwnerDBName *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	// The status of initial schema synchronization. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Migrating**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetConstraints() []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	return s.Constraints
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetObjectDefinition() *string {
	return s.ObjectDefinition
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetObjectName() *string {
	return s.ObjectName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetConstraints(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Constraints = v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) Validate() error {
	if s.Constraints != nil {
		for _, item := range s.Constraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints struct {
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	// The error message returned if constraints failed to be created.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The syntax to create constraints.
	//
	// example:
	//
	// CREATE SEQUENCE "public"."collections_id_seq"   MINVALUE 1   MAXVALUE 9223372036854775807   START 249   INCREMENT BY 1 ;
	ObjectDefinition *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// customer
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The type of the object. Valid value: **Table**.
	//
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	SourceOwnerDBName *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	// The status of constraint creation. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Migrating**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetDestinationOwnerDBName() *string {
	return s.DestinationOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetObjectDefinition() *string {
	return s.ObjectDefinition
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetObjectName() *string {
	return s.ObjectName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetSourceOwnerDBName() *string {
	return s.SourceOwnerDBName
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GetStatus() *string {
	return s.Status
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) Validate() error {
	return dara.Validate(s)
}
