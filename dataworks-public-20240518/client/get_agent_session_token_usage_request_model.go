// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionTokenUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionTokenUsageRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionTokenUsageRequest
	GetJsonrpc() *string
	SetParams(v *GetAgentSessionTokenUsageRequestParams) *GetAgentSessionTokenUsageRequest
	GetParams() *GetAgentSessionTokenUsageRequestParams
}

type GetAgentSessionTokenUsageRequest struct {
	// example:
	//
	// 1033814166
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                 `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *GetAgentSessionTokenUsageRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s GetAgentSessionTokenUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionTokenUsageRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionTokenUsageRequest) GetParams() *GetAgentSessionTokenUsageRequestParams {
	return s.Params
}

func (s *GetAgentSessionTokenUsageRequest) SetId(v string) *GetAgentSessionTokenUsageRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionTokenUsageRequest) SetJsonrpc(v string) *GetAgentSessionTokenUsageRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionTokenUsageRequest) SetParams(v *GetAgentSessionTokenUsageRequestParams) *GetAgentSessionTokenUsageRequest {
	s.Params = v
	return s
}

func (s *GetAgentSessionTokenUsageRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionTokenUsageRequestParams struct {
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetAgentSessionTokenUsageRequestParams) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageRequestParams) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAgentSessionTokenUsageRequestParams) SetSessionId(v string) *GetAgentSessionTokenUsageRequestParams {
	s.SessionId = &v
	return s
}

func (s *GetAgentSessionTokenUsageRequestParams) Validate() error {
	return dara.Validate(s)
}
