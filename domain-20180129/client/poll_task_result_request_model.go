// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *PollTaskResultRequest
	GetDomainName() *string
	SetInstanceId(v string) *PollTaskResultRequest
	GetInstanceId() *string
	SetLang(v string) *PollTaskResultRequest
	GetLang() *string
	SetPageNum(v int32) *PollTaskResultRequest
	GetPageNum() *int32
	SetPageSize(v int32) *PollTaskResultRequest
	GetPageSize() *int32
	SetTaskNo(v string) *PollTaskResultRequest
	GetTaskNo() *string
	SetTaskResultStatus(v int32) *PollTaskResultRequest
	GetTaskResultStatus() *int32
	SetUserClientIp(v string) *PollTaskResultRequest
	GetUserClientIp() *string
}

type PollTaskResultRequest struct {
	// Domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Domain instance ID.
	//
	// The system automatically generates this after the information template is created successfully. You can invoke the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) API to query the information template ID.
	//
	// example:
	//
	// S20181T0WLI85212
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value is **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size. Maximum value is **1000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Job number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// Task result status. Valid values:
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// example:
	//
	// 2
	TaskResultStatus *int32 `json:"TaskResultStatus,omitempty" xml:"TaskResultStatus,omitempty"`
	// User IP address. It can be set to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s PollTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s PollTaskResultRequest) GoString() string {
	return s.String()
}

func (s *PollTaskResultRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PollTaskResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollTaskResultRequest) GetLang() *string {
	return s.Lang
}

func (s *PollTaskResultRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *PollTaskResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PollTaskResultRequest) GetTaskNo() *string {
	return s.TaskNo
}

func (s *PollTaskResultRequest) GetTaskResultStatus() *int32 {
	return s.TaskResultStatus
}

func (s *PollTaskResultRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *PollTaskResultRequest) SetDomainName(v string) *PollTaskResultRequest {
	s.DomainName = &v
	return s
}

func (s *PollTaskResultRequest) SetInstanceId(v string) *PollTaskResultRequest {
	s.InstanceId = &v
	return s
}

func (s *PollTaskResultRequest) SetLang(v string) *PollTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *PollTaskResultRequest) SetPageNum(v int32) *PollTaskResultRequest {
	s.PageNum = &v
	return s
}

func (s *PollTaskResultRequest) SetPageSize(v int32) *PollTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *PollTaskResultRequest) SetTaskNo(v string) *PollTaskResultRequest {
	s.TaskNo = &v
	return s
}

func (s *PollTaskResultRequest) SetTaskResultStatus(v int32) *PollTaskResultRequest {
	s.TaskResultStatus = &v
	return s
}

func (s *PollTaskResultRequest) SetUserClientIp(v string) *PollTaskResultRequest {
	s.UserClientIp = &v
	return s
}

func (s *PollTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
