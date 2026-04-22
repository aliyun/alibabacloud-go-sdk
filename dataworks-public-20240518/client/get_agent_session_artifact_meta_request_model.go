// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionArtifactMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionArtifactMetaRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionArtifactMetaRequest
	GetJsonrpc() *string
	SetParams(v *GetAgentSessionArtifactMetaRequestParams) *GetAgentSessionArtifactMetaRequest
	GetParams() *GetAgentSessionArtifactMetaRequestParams
}

type GetAgentSessionArtifactMetaRequest struct {
	// example:
	//
	// 900335678024
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                   `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *GetAgentSessionArtifactMetaRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s GetAgentSessionArtifactMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionArtifactMetaRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionArtifactMetaRequest) GetParams() *GetAgentSessionArtifactMetaRequestParams {
	return s.Params
}

func (s *GetAgentSessionArtifactMetaRequest) SetId(v string) *GetAgentSessionArtifactMetaRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionArtifactMetaRequest) SetJsonrpc(v string) *GetAgentSessionArtifactMetaRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionArtifactMetaRequest) SetParams(v *GetAgentSessionArtifactMetaRequestParams) *GetAgentSessionArtifactMetaRequest {
	s.Params = v
	return s
}

func (s *GetAgentSessionArtifactMetaRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionArtifactMetaRequestParams struct {
	// example:
	//
	// mock/mock_report.md
	ArtifactPath *string `json:"ArtifactPath,omitempty" xml:"ArtifactPath,omitempty"`
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetAgentSessionArtifactMetaRequestParams) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaRequestParams) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaRequestParams) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *GetAgentSessionArtifactMetaRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAgentSessionArtifactMetaRequestParams) SetArtifactPath(v string) *GetAgentSessionArtifactMetaRequestParams {
	s.ArtifactPath = &v
	return s
}

func (s *GetAgentSessionArtifactMetaRequestParams) SetSessionId(v string) *GetAgentSessionArtifactMetaRequestParams {
	s.SessionId = &v
	return s
}

func (s *GetAgentSessionArtifactMetaRequestParams) Validate() error {
	return dara.Validate(s)
}
