// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSessionTitleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateAgentSessionTitleRequest
	GetId() *string
	SetJsonrpc(v string) *UpdateAgentSessionTitleRequest
	GetJsonrpc() *string
	SetParams(v *UpdateAgentSessionTitleRequestParams) *UpdateAgentSessionTitleRequest
	GetParams() *UpdateAgentSessionTitleRequestParams
}

type UpdateAgentSessionTitleRequest struct {
	// example:
	//
	// updateagentsessiontitle-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                               `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *UpdateAgentSessionTitleRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s UpdateAgentSessionTitleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleRequest) GetId() *string {
	return s.Id
}

func (s *UpdateAgentSessionTitleRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UpdateAgentSessionTitleRequest) GetParams() *UpdateAgentSessionTitleRequestParams {
	return s.Params
}

func (s *UpdateAgentSessionTitleRequest) SetId(v string) *UpdateAgentSessionTitleRequest {
	s.Id = &v
	return s
}

func (s *UpdateAgentSessionTitleRequest) SetJsonrpc(v string) *UpdateAgentSessionTitleRequest {
	s.Jsonrpc = &v
	return s
}

func (s *UpdateAgentSessionTitleRequest) SetParams(v *UpdateAgentSessionTitleRequestParams) *UpdateAgentSessionTitleRequest {
	s.Params = v
	return s
}

func (s *UpdateAgentSessionTitleRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentSessionTitleRequestParams struct {
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-42d3-a456-426614174000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// New title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateAgentSessionTitleRequestParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateAgentSessionTitleRequestParams) GetTitle() *string {
	return s.Title
}

func (s *UpdateAgentSessionTitleRequestParams) SetSessionId(v string) *UpdateAgentSessionTitleRequestParams {
	s.SessionId = &v
	return s
}

func (s *UpdateAgentSessionTitleRequestParams) SetTitle(v string) *UpdateAgentSessionTitleRequestParams {
	s.Title = &v
	return s
}

func (s *UpdateAgentSessionTitleRequestParams) Validate() error {
	return dara.Validate(s)
}
