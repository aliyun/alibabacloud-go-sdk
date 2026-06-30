// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetOperationRecordsRequest
	GetAppType() *string
	SetLanguage(v string) *GetOperationRecordsRequest
	GetLanguage() *string
	SetProcessInstanceId(v string) *GetOperationRecordsRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *GetOperationRecordsRequest
	GetSystemToken() *string
}

type GetOperationRecordsRequest struct {
	AppType           *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	SystemToken       *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetOperationRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetOperationRecordsRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetOperationRecordsRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetOperationRecordsRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetOperationRecordsRequest) SetAppType(v string) *GetOperationRecordsRequest {
	s.AppType = &v
	return s
}

func (s *GetOperationRecordsRequest) SetLanguage(v string) *GetOperationRecordsRequest {
	s.Language = &v
	return s
}

func (s *GetOperationRecordsRequest) SetProcessInstanceId(v string) *GetOperationRecordsRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetOperationRecordsRequest) SetSystemToken(v string) *GetOperationRecordsRequest {
	s.SystemToken = &v
	return s
}

func (s *GetOperationRecordsRequest) Validate() error {
	return dara.Validate(s)
}
