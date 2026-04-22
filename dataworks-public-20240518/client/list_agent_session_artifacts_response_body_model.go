// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) *ListAgentSessionArtifactsResponseBody
	GetJsonRpcResponse() *ListAgentSessionArtifactsResponseBodyJsonRpcResponse
	SetRequestId(v string) *ListAgentSessionArtifactsResponseBody
	GetRequestId() *string
}

type ListAgentSessionArtifactsResponseBody struct {
	JsonRpcResponse *ListAgentSessionArtifactsResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentSessionArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsResponseBody) GetJsonRpcResponse() *ListAgentSessionArtifactsResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ListAgentSessionArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSessionArtifactsResponseBody) SetJsonRpcResponse(v *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) *ListAgentSessionArtifactsResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ListAgentSessionArtifactsResponseBody) SetRequestId(v string) *ListAgentSessionArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionArtifactsResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 28477817
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                                     `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) GetResult() *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) SetId(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) SetResult(v *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) *ListAgentSessionArtifactsResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult struct {
	Artifacts []*ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// example:
	//
	// 29
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESExFsbyH...
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) GetArtifacts() []*ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts {
	return s.Artifacts
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) SetArtifacts(v []*ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult {
	s.Artifacts = v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) SetMaxResults(v int32) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) SetNextToken(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResult) Validate() error {
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts struct {
	// example:
	//
	// mock_report.md
	ArtifactName *string `json:"ArtifactName,omitempty" xml:"ArtifactName,omitempty"`
	// example:
	//
	// mock/mock_report.md
	ArtifactPath *string `json:"ArtifactPath,omitempty" xml:"ArtifactPath,omitempty"`
	// example:
	//
	// md
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) GetArtifactName() *string {
	return s.ArtifactName
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) SetArtifactName(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts {
	s.ArtifactName = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) SetArtifactPath(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts {
	s.ArtifactPath = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) SetArtifactType(v string) *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListAgentSessionArtifactsResponseBodyJsonRpcResponseResultArtifacts) Validate() error {
	return dara.Validate(s)
}
