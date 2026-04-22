// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionArtifactMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) *GetAgentSessionArtifactMetaResponseBody
	GetJsonRpcResponse() *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse
	SetRequestId(v string) *GetAgentSessionArtifactMetaResponseBody
	GetRequestId() *string
}

type GetAgentSessionArtifactMetaResponseBody struct {
	JsonRpcResponse *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// example:
	//
	// CE70C54F-A3BD-5C19-88EF-2A7D3451C449
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentSessionArtifactMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaResponseBody) GetJsonRpcResponse() *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *GetAgentSessionArtifactMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSessionArtifactMetaResponseBody) SetJsonRpcResponse(v *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) *GetAgentSessionArtifactMetaResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBody) SetRequestId(v string) *GetAgentSessionArtifactMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 300010555
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                                       `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) GetResult() *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) SetId(v string) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) SetJsonrpc(v string) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) SetResult(v *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult struct {
	ArtifactContent *string `json:"ArtifactContent,omitempty" xml:"ArtifactContent,omitempty"`
	// example:
	//
	// mock_report.md
	ArtifactName *string `json:"ArtifactName,omitempty" xml:"ArtifactName,omitempty"`
	// example:
	//
	// mock/mock_report.md
	ArtifactPath *string `json:"ArtifactPath,omitempty" xml:"ArtifactPath,omitempty"`
}

func (s GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) GetArtifactContent() *string {
	return s.ArtifactContent
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) GetArtifactName() *string {
	return s.ArtifactName
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) SetArtifactContent(v string) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult {
	s.ArtifactContent = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) SetArtifactName(v string) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult {
	s.ArtifactName = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) SetArtifactPath(v string) *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult {
	s.ArtifactPath = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
