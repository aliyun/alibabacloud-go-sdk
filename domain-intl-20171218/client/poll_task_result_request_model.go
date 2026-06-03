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
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskNo           *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskResultStatus *int32  `json:"TaskResultStatus,omitempty" xml:"TaskResultStatus,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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
