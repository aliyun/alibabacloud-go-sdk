// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryTaskDetailListRequest
	GetDomainName() *string
	SetInstanceId(v string) *QueryTaskDetailListRequest
	GetInstanceId() *string
	SetLang(v string) *QueryTaskDetailListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryTaskDetailListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryTaskDetailListRequest
	GetPageSize() *int32
	SetTaskNo(v string) *QueryTaskDetailListRequest
	GetTaskNo() *string
	SetTaskStatus(v int32) *QueryTaskDetailListRequest
	GetTaskStatus() *int32
	SetUserClientIp(v string) *QueryTaskDetailListRequest
	GetUserClientIp() *string
}

type QueryTaskDetailListRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The instance ID of the domain name.
	//
	// example:
	//
	// S20179H1BBI9test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the error message to return if the request fails. Valid value:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Maximum value: **1000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// The task status. Valid value:
	//
	// 	- **0**: waiting for execution
	//
	// 	- **1**: being executed
	//
	// 	- **2**: successful
	//
	// 	- **3**: failed
	//
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The IP address of the client. Set the value to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTaskDetailListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryTaskDetailListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskDetailListRequest) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailListRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryTaskDetailListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTaskDetailListRequest) SetDomainName(v string) *QueryTaskDetailListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetInstanceId(v string) *QueryTaskDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetLang(v string) *QueryTaskDetailListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetPageNum(v int32) *QueryTaskDetailListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetPageSize(v int32) *QueryTaskDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetTaskNo(v string) *QueryTaskDetailListRequest {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetTaskStatus(v int32) *QueryTaskDetailListRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetUserClientIp(v string) *QueryTaskDetailListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTaskDetailListRequest) Validate() error {
	return dara.Validate(s)
}
