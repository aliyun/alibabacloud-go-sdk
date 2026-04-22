// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentSessionArtifactsRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentSessionArtifactsRequest
	GetJsonrpc() *string
	SetParams(v *ListAgentSessionArtifactsRequestParams) *ListAgentSessionArtifactsRequest
	GetParams() *ListAgentSessionArtifactsRequestParams
}

type ListAgentSessionArtifactsRequest struct {
	// example:
	//
	// 10001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                 `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *ListAgentSessionArtifactsRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s ListAgentSessionArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionArtifactsRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionArtifactsRequest) GetParams() *ListAgentSessionArtifactsRequestParams {
	return s.Params
}

func (s *ListAgentSessionArtifactsRequest) SetId(v string) *ListAgentSessionArtifactsRequest {
	s.Id = &v
	return s
}

func (s *ListAgentSessionArtifactsRequest) SetJsonrpc(v string) *ListAgentSessionArtifactsRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionArtifactsRequest) SetParams(v *ListAgentSessionArtifactsRequestParams) *ListAgentSessionArtifactsRequest {
	s.Params = v
	return s
}

func (s *ListAgentSessionArtifactsRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionArtifactsRequestParams struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESExFsbyH...
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// req_20260421_001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListAgentSessionArtifactsRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsRequestParams) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsRequestParams) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionArtifactsRequestParams) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionArtifactsRequestParams) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSessionArtifactsRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAgentSessionArtifactsRequestParams) SetMaxResults(v int32) *ListAgentSessionArtifactsRequestParams {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionArtifactsRequestParams) SetNextToken(v string) *ListAgentSessionArtifactsRequestParams {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionArtifactsRequestParams) SetRequestId(v string) *ListAgentSessionArtifactsRequestParams {
	s.RequestId = &v
	return s
}

func (s *ListAgentSessionArtifactsRequestParams) SetSessionId(v string) *ListAgentSessionArtifactsRequestParams {
	s.SessionId = &v
	return s
}

func (s *ListAgentSessionArtifactsRequestParams) Validate() error {
	return dara.Validate(s)
}
