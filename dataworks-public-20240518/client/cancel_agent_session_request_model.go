// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CancelAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *CancelAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v *CancelAgentSessionRequestParams) *CancelAgentSessionRequest
	GetParams() *CancelAgentSessionRequestParams
}

type CancelAgentSessionRequest struct {
	// example:
	//
	// 676303114031776
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                          `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *CancelAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s CancelAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *CancelAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CancelAgentSessionRequest) GetParams() *CancelAgentSessionRequestParams {
	return s.Params
}

func (s *CancelAgentSessionRequest) SetId(v string) *CancelAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *CancelAgentSessionRequest) SetJsonrpc(v string) *CancelAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *CancelAgentSessionRequest) SetParams(v *CancelAgentSessionRequestParams) *CancelAgentSessionRequest {
	s.Params = v
	return s
}

func (s *CancelAgentSessionRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelAgentSessionRequestParams struct {
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CancelAgentSessionRequestParams) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionRequestParams) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *CancelAgentSessionRequestParams) SetSessionId(v string) *CancelAgentSessionRequestParams {
	s.SessionId = &v
	return s
}

func (s *CancelAgentSessionRequestParams) Validate() error {
	return dara.Validate(s)
}
