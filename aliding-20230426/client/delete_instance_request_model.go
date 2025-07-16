// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *DeleteInstanceRequest
	GetAppType() *string
	SetLanguage(v string) *DeleteInstanceRequest
	GetLanguage() *string
	SetProcessInstanceId(v string) *DeleteInstanceRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *DeleteInstanceRequest
	GetSystemToken() *string
}

type DeleteInstanceRequest struct {
	// example:
	//
	// APP_PBKTxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// f30233fb-72xxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetAppType() *string {
	return s.AppType
}

func (s *DeleteInstanceRequest) GetLanguage() *string {
	return s.Language
}

func (s *DeleteInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *DeleteInstanceRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *DeleteInstanceRequest) SetAppType(v string) *DeleteInstanceRequest {
	s.AppType = &v
	return s
}

func (s *DeleteInstanceRequest) SetLanguage(v string) *DeleteInstanceRequest {
	s.Language = &v
	return s
}

func (s *DeleteInstanceRequest) SetProcessInstanceId(v string) *DeleteInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetSystemToken(v string) *DeleteInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
