// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetRunningTasksRequest
	GetAppType() *string
	SetLanguage(v string) *GetRunningTasksRequest
	GetLanguage() *string
	SetProcessCodes(v string) *GetRunningTasksRequest
	GetProcessCodes() *string
	SetProcessInstanceId(v string) *GetRunningTasksRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *GetRunningTasksRequest
	GetSystemToken() *string
}

type GetRunningTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ["xx","xxx"]
	ProcessCodes *string `json:"ProcessCodes,omitempty" xml:"ProcessCodes,omitempty"`
	// example:
	//
	// instxxxxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetRunningTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksRequest) GoString() string {
	return s.String()
}

func (s *GetRunningTasksRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetRunningTasksRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetRunningTasksRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetRunningTasksRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetRunningTasksRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetRunningTasksRequest) SetAppType(v string) *GetRunningTasksRequest {
	s.AppType = &v
	return s
}

func (s *GetRunningTasksRequest) SetLanguage(v string) *GetRunningTasksRequest {
	s.Language = &v
	return s
}

func (s *GetRunningTasksRequest) SetProcessCodes(v string) *GetRunningTasksRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetRunningTasksRequest) SetProcessInstanceId(v string) *GetRunningTasksRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksRequest) SetSystemToken(v string) *GetRunningTasksRequest {
	s.SystemToken = &v
	return s
}

func (s *GetRunningTasksRequest) Validate() error {
	return dara.Validate(s)
}
