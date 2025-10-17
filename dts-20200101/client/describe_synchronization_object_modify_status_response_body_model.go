// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationObjectModifyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetDataInitializationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus
	SetDataSynchronizationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetDataSynchronizationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus
	SetErrCode(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetErrMessage() *string
	SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetErrorMessage() *string
	SetPrecheckStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetPrecheckStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus
	SetRequestId(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetStatus() *string
	SetStructureInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetStructureInitializationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus
	SetSuccess(v string) *DescribeSynchronizationObjectModifyStatusResponseBody
	GetSuccess() *string
}

type DescribeSynchronizationObjectModifyStatusResponseBody struct {
	// The status of full data synchronization.
	DataInitializationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data synchronization.
	//
	// >  This parameter and its sub-parameters will be removed in the future.
	DataSynchronizationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
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
	// The error message returned if the task failed to modify the objects to be synchronized.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck status.
	PrecheckStatus *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B38C644B-4395-4F6F-86E3-592F26BE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task that changes the objects to be synchronized. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Migrating**: The task is running.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of schema synchronization.
	StructureInitializationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetDataInitializationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetDataSynchronizationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetPrecheckStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetStructureInitializationStatus() *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) Validate() error {
	if s.DataInitializationStatus != nil {
		if err := s.DataInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronizationStatus != nil {
		if err := s.DataSynchronizationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.PrecheckStatus != nil {
		if err := s.PrecheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitializationStatus != nil {
		if err := s.StructureInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus struct {
	// The error message returned if full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of full data synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been synchronized during full data synchronization.
	//
	// example:
	//
	// 39754
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of full data synchronization. Valid values:
	//
	// 	- **NotStarted**: Full data synchronization is not started.
	//
	// 	- **Migrating**: Full data synchronization is in progress.
	//
	// 	- **Failed**: Full data synchronization failed.
	//
	// 	- **Finished**: Full data synchronization is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus struct {
	// The synchronization latency, in seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The error message returned if incremental data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of incremental data synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of incremental data synchronization. Valid values:
	//
	// 	- **NotStarted**: Incremental data synchronization is not started.
	//
	// 	- **Migrating**: Incremental data synchronization is in progress.
	//
	// 	- **Failed**: Incremental data synchronization failed.
	//
	// 	- **Finished**: Incremental data synchronization is completed.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GetDelay() *string {
	return s.Delay
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The precheck progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck status.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GetDetail() []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail struct {
	// The precheck result. Valid values:
	//
	// 	- Success: The task passed the precheck.
	//
	// 	- Failed: The task failed to pass the precheck.
	//
	// example:
	//
	// Success
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.xxx.xx\\" (using password: YES)
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The method to fix the precheck failure.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is Failed.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus struct {
	// The error message returned if schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: Table \\"customer\\" already exists
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables whose schemas have been synchronized.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of schema synchronization. Valid values:
	//
	// 	- **NotStarted**: Schema synchronization is not started.
	//
	// 	- **Migrating**: Schema synchronization is in progress.
	//
	// 	- **Failed**: Schema synchronization failed.
	//
	// 	- **Finished**: Schema synchronization is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}
