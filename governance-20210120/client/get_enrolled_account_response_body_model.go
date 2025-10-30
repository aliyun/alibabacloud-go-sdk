// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnrolledAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountUid(v int64) *GetEnrolledAccountResponseBody
	GetAccountUid() *int64
	SetBaselineId(v string) *GetEnrolledAccountResponseBody
	GetBaselineId() *string
	SetBaselineItems(v []*GetEnrolledAccountResponseBodyBaselineItems) *GetEnrolledAccountResponseBody
	GetBaselineItems() []*GetEnrolledAccountResponseBodyBaselineItems
	SetCreateTime(v string) *GetEnrolledAccountResponseBody
	GetCreateTime() *string
	SetDisplayName(v string) *GetEnrolledAccountResponseBody
	GetDisplayName() *string
	SetErrorInfo(v *GetEnrolledAccountResponseBodyErrorInfo) *GetEnrolledAccountResponseBody
	GetErrorInfo() *GetEnrolledAccountResponseBodyErrorInfo
	SetFolderId(v string) *GetEnrolledAccountResponseBody
	GetFolderId() *string
	SetInitialized(v bool) *GetEnrolledAccountResponseBody
	GetInitialized() *bool
	SetInputs(v *GetEnrolledAccountResponseBodyInputs) *GetEnrolledAccountResponseBody
	GetInputs() *GetEnrolledAccountResponseBodyInputs
	SetMasterAccountUid(v int64) *GetEnrolledAccountResponseBody
	GetMasterAccountUid() *int64
	SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBody
	GetPayerAccountUid() *int64
	SetProgress(v []*GetEnrolledAccountResponseBodyProgress) *GetEnrolledAccountResponseBody
	GetProgress() []*GetEnrolledAccountResponseBodyProgress
	SetRequestId(v string) *GetEnrolledAccountResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetEnrolledAccountResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetEnrolledAccountResponseBody
	GetUpdateTime() *string
}

type GetEnrolledAccountResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// 12868156179*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The ID of the baseline that is implemented.
	//
	// example:
	//
	// afb-bp1adadfadsf***
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The array that contains baseline items.
	BaselineItems []*GetEnrolledAccountResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The time when the account was created.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// test-account
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The error message.
	//
	// >  This parameter is returned if the value of `Status` is `Failed` or `ScheduleFailed`.
	ErrorInfo *GetEnrolledAccountResponseBodyErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Indicates whether the initialization is complete. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	Initialized *bool `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// Input parameters used to create an account.
	Inputs *GetEnrolledAccountResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The ID of the management account of the resource directory to which the account belongs.
	//
	// example:
	//
	// 19534534552*****
	MasterAccountUid *int64 `json:"MasterAccountUid,omitempty" xml:"MasterAccountUid,omitempty"`
	// The ID of the settlement account.
	//
	// example:
	//
	// 19534534552*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The progress of the applying the baseline to the account.
	Progress []*GetEnrolledAccountResponseBodyProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- Pending: The account is pending to be created.
	//
	// 	- Running: The account is being created.
	//
	// 	- Finished: The account is created.
	//
	// 	- Failed: The account fails to be created.
	//
	// 	- Scheduling: The account is being scheduled.
	//
	// 	- ScheduleFailed: The account fails to be scheduled.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetEnrolledAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBody) GetAccountUid() *int64 {
	return s.AccountUid
}

func (s *GetEnrolledAccountResponseBody) GetBaselineId() *string {
	return s.BaselineId
}

func (s *GetEnrolledAccountResponseBody) GetBaselineItems() []*GetEnrolledAccountResponseBodyBaselineItems {
	return s.BaselineItems
}

func (s *GetEnrolledAccountResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEnrolledAccountResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEnrolledAccountResponseBody) GetErrorInfo() *GetEnrolledAccountResponseBodyErrorInfo {
	return s.ErrorInfo
}

func (s *GetEnrolledAccountResponseBody) GetFolderId() *string {
	return s.FolderId
}

func (s *GetEnrolledAccountResponseBody) GetInitialized() *bool {
	return s.Initialized
}

func (s *GetEnrolledAccountResponseBody) GetInputs() *GetEnrolledAccountResponseBodyInputs {
	return s.Inputs
}

func (s *GetEnrolledAccountResponseBody) GetMasterAccountUid() *int64 {
	return s.MasterAccountUid
}

func (s *GetEnrolledAccountResponseBody) GetPayerAccountUid() *int64 {
	return s.PayerAccountUid
}

func (s *GetEnrolledAccountResponseBody) GetProgress() []*GetEnrolledAccountResponseBodyProgress {
	return s.Progress
}

func (s *GetEnrolledAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEnrolledAccountResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEnrolledAccountResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetEnrolledAccountResponseBody) SetAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetBaselineId(v string) *GetEnrolledAccountResponseBody {
	s.BaselineId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetBaselineItems(v []*GetEnrolledAccountResponseBodyBaselineItems) *GetEnrolledAccountResponseBody {
	s.BaselineItems = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetCreateTime(v string) *GetEnrolledAccountResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetDisplayName(v string) *GetEnrolledAccountResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetErrorInfo(v *GetEnrolledAccountResponseBodyErrorInfo) *GetEnrolledAccountResponseBody {
	s.ErrorInfo = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetFolderId(v string) *GetEnrolledAccountResponseBody {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInitialized(v bool) *GetEnrolledAccountResponseBody {
	s.Initialized = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInputs(v *GetEnrolledAccountResponseBodyInputs) *GetEnrolledAccountResponseBody {
	s.Inputs = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetMasterAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.MasterAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.PayerAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetProgress(v []*GetEnrolledAccountResponseBodyProgress) *GetEnrolledAccountResponseBody {
	s.Progress = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetRequestId(v string) *GetEnrolledAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetStatus(v string) *GetEnrolledAccountResponseBody {
	s.Status = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetUpdateTime(v string) *GetEnrolledAccountResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) Validate() error {
	if s.BaselineItems != nil {
		for _, item := range s.BaselineItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ErrorInfo != nil {
		if err := s.ErrorInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Inputs != nil {
		if err := s.Inputs.Validate(); err != nil {
			return err
		}
	}
	if s.Progress != nil {
		for _, item := range s.Progress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEnrolledAccountResponseBodyBaselineItems struct {
	// The configuration of the baseline item.
	//
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether baseline item is skipped. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetEnrolledAccountResponseBodyBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) GetName() *string {
	return s.Name
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) GetSkip() *bool {
	return s.Skip
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetConfig(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Config = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetName(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetSkip(v bool) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Skip = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetVersion(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Version = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) Validate() error {
	return dara.Validate(s)
}

type GetEnrolledAccountResponseBodyErrorInfo struct {
	// The error code.
	//
	// example:
	//
	// CompliancePackExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The compliance pack already exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The recommended solution.
	//
	// example:
	//
	// https://next.api.aliyun.com/troubleshoot?q=CompliancePackExists\\\\u0026product=Config
	Recommend *string `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D5EAA86-2D41-5CB7-8DA7-B60093ACAA4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEnrolledAccountResponseBodyErrorInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyErrorInfo) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) GetCode() *string {
	return s.Code
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) GetMessage() *string {
	return s.Message
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) GetRecommend() *string {
	return s.Recommend
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetCode(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Code = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetMessage(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Message = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRecommend(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Recommend = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRequestId(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.RequestId = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) Validate() error {
	return dara.Validate(s)
}

type GetEnrolledAccountResponseBodyInputs struct {
	// The prefix of the account name.
	//
	// example:
	//
	// test-account
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 12868156179*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The baseline items.
	BaselineItems []*GetEnrolledAccountResponseBodyInputsBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The display name of the account.
	//
	// example:
	//
	// test-account
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the settlement account.
	//
	// example:
	//
	// 19534534552*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The tag.
	Tag []*GetEnrolledAccountResponseBodyInputsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetEnrolledAccountResponseBodyInputs) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputs) GetAccountNamePrefix() *string {
	return s.AccountNamePrefix
}

func (s *GetEnrolledAccountResponseBodyInputs) GetAccountUid() *int64 {
	return s.AccountUid
}

func (s *GetEnrolledAccountResponseBodyInputs) GetBaselineItems() []*GetEnrolledAccountResponseBodyInputsBaselineItems {
	return s.BaselineItems
}

func (s *GetEnrolledAccountResponseBodyInputs) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEnrolledAccountResponseBodyInputs) GetFolderId() *string {
	return s.FolderId
}

func (s *GetEnrolledAccountResponseBodyInputs) GetPayerAccountUid() *int64 {
	return s.PayerAccountUid
}

func (s *GetEnrolledAccountResponseBodyInputs) GetTag() []*GetEnrolledAccountResponseBodyInputsTag {
	return s.Tag
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountNamePrefix(v string) *GetEnrolledAccountResponseBodyInputs {
	s.AccountNamePrefix = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetBaselineItems(v []*GetEnrolledAccountResponseBodyInputsBaselineItems) *GetEnrolledAccountResponseBodyInputs {
	s.BaselineItems = v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetDisplayName(v string) *GetEnrolledAccountResponseBodyInputs {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetFolderId(v string) *GetEnrolledAccountResponseBodyInputs {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.PayerAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetTag(v []*GetEnrolledAccountResponseBodyInputsTag) *GetEnrolledAccountResponseBodyInputs {
	s.Tag = v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) Validate() error {
	if s.BaselineItems != nil {
		for _, item := range s.BaselineItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEnrolledAccountResponseBodyInputsBaselineItems struct {
	// The configurations of the baseline item.
	//
	// example:
	//
	// {\\"Contacts\\":[{\\"Name\\":\\"governance\\",\\"Email\\":\\"wibud****@gmail.com\\",\\"Mobile\\":\\"1234\\",\\"Position\\":\\"Other\\"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether baseline item is skipped. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) GetName() *string {
	return s.Name
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) GetSkip() *bool {
	return s.Skip
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetConfig(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Config = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetName(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetSkip(v bool) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Skip = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetVersion(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Version = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) Validate() error {
	return dara.Validate(s)
}

type GetEnrolledAccountResponseBodyInputsTag struct {
	// The tag key.
	//
	// example:
	//
	// product
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// governance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputsTag) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputsTag) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputsTag) GetKey() *string {
	return s.Key
}

func (s *GetEnrolledAccountResponseBodyInputsTag) GetValue() *string {
	return s.Value
}

func (s *GetEnrolledAccountResponseBodyInputsTag) SetKey(v string) *GetEnrolledAccountResponseBodyInputsTag {
	s.Key = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsTag) SetValue(v string) *GetEnrolledAccountResponseBodyInputsTag {
	s.Value = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsTag) Validate() error {
	return dara.Validate(s)
}

type GetEnrolledAccountResponseBodyProgress struct {
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of applying the baseline to the account. Valid values:
	//
	// 	- Pending: The baseline is pending to be applied to the account.
	//
	// 	- Running: The baseline is being applied to the account.
	//
	// 	- Finished: : The baseline is applied to the account.
	//
	// 	- Failed: : The baseline fails to be applied to the account.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEnrolledAccountResponseBodyProgress) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyProgress) GetName() *string {
	return s.Name
}

func (s *GetEnrolledAccountResponseBodyProgress) GetStatus() *string {
	return s.Status
}

func (s *GetEnrolledAccountResponseBodyProgress) SetName(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyProgress) SetStatus(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Status = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyProgress) Validate() error {
	return dara.Validate(s)
}
