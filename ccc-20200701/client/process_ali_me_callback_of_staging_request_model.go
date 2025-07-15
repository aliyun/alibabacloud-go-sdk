// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessAliMeCallbackOfStagingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ProcessAliMeCallbackOfStagingRequest
	GetData() *string
	SetToken(v string) *ProcessAliMeCallbackOfStagingRequest
	GetToken() *string
}

type ProcessAliMeCallbackOfStagingRequest struct {
	Data  *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ProcessAliMeCallbackOfStagingRequest) String() string {
	return dara.Prettify(s)
}

func (s ProcessAliMeCallbackOfStagingRequest) GoString() string {
	return s.String()
}

func (s *ProcessAliMeCallbackOfStagingRequest) GetData() *string {
	return s.Data
}

func (s *ProcessAliMeCallbackOfStagingRequest) GetToken() *string {
	return s.Token
}

func (s *ProcessAliMeCallbackOfStagingRequest) SetData(v string) *ProcessAliMeCallbackOfStagingRequest {
	s.Data = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingRequest) SetToken(v string) *ProcessAliMeCallbackOfStagingRequest {
	s.Token = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingRequest) Validate() error {
	return dara.Validate(s)
}
