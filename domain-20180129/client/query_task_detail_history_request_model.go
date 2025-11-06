// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryTaskDetailHistoryRequest
	GetDomainName() *string
	SetDomainNameCursor(v string) *QueryTaskDetailHistoryRequest
	GetDomainNameCursor() *string
	SetLang(v string) *QueryTaskDetailHistoryRequest
	GetLang() *string
	SetPageSize(v int32) *QueryTaskDetailHistoryRequest
	GetPageSize() *int32
	SetTaskDetailNoCursor(v string) *QueryTaskDetailHistoryRequest
	GetTaskDetailNoCursor() *string
	SetTaskNo(v string) *QueryTaskDetailHistoryRequest
	GetTaskNo() *string
	SetTaskStatus(v int32) *QueryTaskDetailHistoryRequest
	GetTaskStatus() *int32
	SetUserClientIp(v string) *QueryTaskDetailHistoryRequest
	GetUserClientIp() *string
}

type QueryTaskDetailHistoryRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// example.com
	DomainNameCursor *string `json:"DomainNameCursor,omitempty" xml:"DomainNameCursor,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 75addb07-28a3-450e-b5ec
	TaskDetailNoCursor *string `json:"TaskDetailNoCursor,omitempty" xml:"TaskDetailNoCursor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskDetailHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailHistoryRequest) GetDomainNameCursor() *string {
	return s.DomainNameCursor
}

func (s *QueryTaskDetailHistoryRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTaskDetailHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskDetailHistoryRequest) GetTaskDetailNoCursor() *string {
	return s.TaskDetailNoCursor
}

func (s *QueryTaskDetailHistoryRequest) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailHistoryRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryTaskDetailHistoryRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTaskDetailHistoryRequest) SetDomainName(v string) *QueryTaskDetailHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetDomainNameCursor(v string) *QueryTaskDetailHistoryRequest {
	s.DomainNameCursor = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetLang(v string) *QueryTaskDetailHistoryRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetPageSize(v int32) *QueryTaskDetailHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskDetailNoCursor(v string) *QueryTaskDetailHistoryRequest {
	s.TaskDetailNoCursor = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskNo(v string) *QueryTaskDetailHistoryRequest {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskStatus(v int32) *QueryTaskDetailHistoryRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetUserClientIp(v string) *QueryTaskDetailHistoryRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) Validate() error {
	return dara.Validate(s)
}
