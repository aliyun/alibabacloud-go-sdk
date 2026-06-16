// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimeVersionsResult
	GetCode() *string
	SetData(v *ListAgentRuntimeVersionsOutput) *ListAgentRuntimeVersionsResult
	GetData() *ListAgentRuntimeVersionsOutput
	SetRequestId(v string) *ListAgentRuntimeVersionsResult
	GetRequestId() *string
}

type ListAgentRuntimeVersionsResult struct {
	// Indicates whether the request succeeded. A value of `SUCCESS` is returned on success. Otherwise, an error type is returned, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Contains the list of agent runtime versions.
	//
	// example:
	//
	// {}
	Data *ListAgentRuntimeVersionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request ID for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimeVersionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimeVersionsResult) GetData() *ListAgentRuntimeVersionsOutput {
	return s.Data
}

func (s *ListAgentRuntimeVersionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimeVersionsResult) SetCode(v string) *ListAgentRuntimeVersionsResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimeVersionsResult) SetData(v *ListAgentRuntimeVersionsOutput) *ListAgentRuntimeVersionsResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimeVersionsResult) SetRequestId(v string) *ListAgentRuntimeVersionsResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimeVersionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
