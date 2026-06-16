// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimeEndpointsResult
	GetCode() *string
	SetData(v *ListAgentRuntimeEndpointsOutput) *ListAgentRuntimeEndpointsResult
	GetData() *ListAgentRuntimeEndpointsOutput
	SetRequestId(v string) *ListAgentRuntimeEndpointsResult
	GetRequestId() *string
}

type ListAgentRuntimeEndpointsResult struct {
	// The status code of the response. `SUCCESS` indicates that the operation succeeded. If the operation fails, an error code is returned, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Detailed information about the agent runtime endpoints.
	//
	// example:
	//
	// {}
	Data *ListAgentRuntimeEndpointsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request identifier for troubleshooting purposes.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimeEndpointsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimeEndpointsResult) GetData() *ListAgentRuntimeEndpointsOutput {
	return s.Data
}

func (s *ListAgentRuntimeEndpointsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimeEndpointsResult) SetCode(v string) *ListAgentRuntimeEndpointsResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) SetData(v *ListAgentRuntimeEndpointsOutput) *ListAgentRuntimeEndpointsResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) SetRequestId(v string) *ListAgentRuntimeEndpointsResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
