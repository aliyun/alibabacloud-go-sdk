// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *LoadAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v *LoadAgentSessionRequestParams) *LoadAgentSessionRequest
	GetParams() *LoadAgentSessionRequestParams
}

type LoadAgentSessionRequest struct {
	// example:
	//
	// 4as3dasf654a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                        `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *LoadAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s LoadAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionRequest) GetParams() *LoadAgentSessionRequestParams {
	return s.Params
}

func (s *LoadAgentSessionRequest) SetId(v string) *LoadAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionRequest) SetJsonrpc(v string) *LoadAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionRequest) SetParams(v *LoadAgentSessionRequestParams) *LoadAgentSessionRequest {
	s.Params = v
	return s
}

func (s *LoadAgentSessionRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoadAgentSessionRequestParams struct {
	Meta *LoadAgentSessionRequestParamsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s LoadAgentSessionRequestParams) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionRequestParams) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionRequestParams) GetMeta() *LoadAgentSessionRequestParamsMeta {
	return s.Meta
}

func (s *LoadAgentSessionRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *LoadAgentSessionRequestParams) SetMeta(v *LoadAgentSessionRequestParamsMeta) *LoadAgentSessionRequestParams {
	s.Meta = v
	return s
}

func (s *LoadAgentSessionRequestParams) SetSessionId(v string) *LoadAgentSessionRequestParams {
	s.SessionId = &v
	return s
}

func (s *LoadAgentSessionRequestParams) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoadAgentSessionRequestParamsMeta struct {
	BeginLogOffset *int64 `json:"BeginLogOffset,omitempty" xml:"BeginLogOffset,omitempty"`
	IsReload       *bool  `json:"IsReload,omitempty" xml:"IsReload,omitempty"`
}

func (s LoadAgentSessionRequestParamsMeta) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionRequestParamsMeta) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionRequestParamsMeta) GetBeginLogOffset() *int64 {
	return s.BeginLogOffset
}

func (s *LoadAgentSessionRequestParamsMeta) GetIsReload() *bool {
	return s.IsReload
}

func (s *LoadAgentSessionRequestParamsMeta) SetBeginLogOffset(v int64) *LoadAgentSessionRequestParamsMeta {
	s.BeginLogOffset = &v
	return s
}

func (s *LoadAgentSessionRequestParamsMeta) SetIsReload(v bool) *LoadAgentSessionRequestParamsMeta {
	s.IsReload = &v
	return s
}

func (s *LoadAgentSessionRequestParamsMeta) Validate() error {
	return dara.Validate(s)
}
