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
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	TaskNo       *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
