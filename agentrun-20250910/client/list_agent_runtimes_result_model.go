// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimesResult
	GetCode() *string
	SetData(v *ListAgentRuntimesOutput) *ListAgentRuntimesResult
	GetData() *ListAgentRuntimesOutput
	SetRequestId(v string) *ListAgentRuntimesResult
	GetRequestId() *string
}

type ListAgentRuntimesResult struct {
	// The status of the request. A value of `SUCCESS` indicates success. Otherwise, this field returns an error type, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of agent runtimes.
	//
	// example:
	//
	// {}
	Data *ListAgentRuntimesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// The unique ID of the request, used for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimesResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimesResult) GetData() *ListAgentRuntimesOutput {
	return s.Data
}

func (s *ListAgentRuntimesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimesResult) SetCode(v string) *ListAgentRuntimesResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimesResult) SetData(v *ListAgentRuntimesOutput) *ListAgentRuntimesResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimesResult) SetRequestId(v string) *ListAgentRuntimesResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
