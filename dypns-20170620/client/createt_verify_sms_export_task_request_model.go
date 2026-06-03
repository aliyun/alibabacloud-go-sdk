// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatetVerifySmsExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *CreatetVerifySmsExportTaskRequest
	GetEndDate() *string
	SetOwnerId(v int64) *CreatetVerifySmsExportTaskRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *CreatetVerifySmsExportTaskRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *CreatetVerifySmsExportTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatetVerifySmsExportTaskRequest
	GetResourceOwnerId() *int64
	SetSchemeName(v string) *CreatetVerifySmsExportTaskRequest
	GetSchemeName() *string
	SetSendStatus(v int64) *CreatetVerifySmsExportTaskRequest
	GetSendStatus() *int64
	SetSignName(v string) *CreatetVerifySmsExportTaskRequest
	GetSignName() *string
	SetStartDate(v string) *CreatetVerifySmsExportTaskRequest
	GetStartDate() *string
	SetTaskName(v string) *CreatetVerifySmsExportTaskRequest
	GetTaskName() *string
	SetTemplateCode(v string) *CreatetVerifySmsExportTaskRequest
	GetTemplateCode() *string
}

type CreatetVerifySmsExportTaskRequest struct {
	// example:
	//
	// 2025-01-08
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 186***9399
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// example:
	//
	// 0
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// example:
	//
	// 2025-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// SMS_483325498
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreatetVerifySmsExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatetVerifySmsExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatetVerifySmsExportTaskRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreatetVerifySmsExportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatetVerifySmsExportTaskRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreatetVerifySmsExportTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatetVerifySmsExportTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatetVerifySmsExportTaskRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CreatetVerifySmsExportTaskRequest) GetSendStatus() *int64 {
	return s.SendStatus
}

func (s *CreatetVerifySmsExportTaskRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreatetVerifySmsExportTaskRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreatetVerifySmsExportTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreatetVerifySmsExportTaskRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreatetVerifySmsExportTaskRequest) SetEndDate(v string) *CreatetVerifySmsExportTaskRequest {
	s.EndDate = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetOwnerId(v int64) *CreatetVerifySmsExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetPhoneNumber(v string) *CreatetVerifySmsExportTaskRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetResourceOwnerAccount(v string) *CreatetVerifySmsExportTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetResourceOwnerId(v int64) *CreatetVerifySmsExportTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetSchemeName(v string) *CreatetVerifySmsExportTaskRequest {
	s.SchemeName = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetSendStatus(v int64) *CreatetVerifySmsExportTaskRequest {
	s.SendStatus = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetSignName(v string) *CreatetVerifySmsExportTaskRequest {
	s.SignName = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetStartDate(v string) *CreatetVerifySmsExportTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetTaskName(v string) *CreatetVerifySmsExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) SetTemplateCode(v string) *CreatetVerifySmsExportTaskRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreatetVerifySmsExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
