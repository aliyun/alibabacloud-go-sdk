// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CancelTaskRequest
	GetLang() *string
	SetTaskNo(v string) *CancelTaskRequest
	GetTaskNo() *string
	SetUserClientIp(v string) *CancelTaskRequest
	GetUserClientIp() *string
}

type CancelTaskRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7f217ae0-61f5-42e2-a1c3-42bad0124****
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelTaskRequest) GetTaskNo() *string {
	return s.TaskNo
}

func (s *CancelTaskRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CancelTaskRequest) SetLang(v string) *CancelTaskRequest {
	s.Lang = &v
	return s
}

func (s *CancelTaskRequest) SetTaskNo(v string) *CancelTaskRequest {
	s.TaskNo = &v
	return s
}

func (s *CancelTaskRequest) SetUserClientIp(v string) *CancelTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *CancelTaskRequest) Validate() error {
	return dara.Validate(s)
}
