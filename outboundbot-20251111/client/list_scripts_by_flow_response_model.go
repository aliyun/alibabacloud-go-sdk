// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptsByFlowResponseBody) *ListScriptsByFlowResponse
	GetBody() *ListScriptsByFlowResponseBody
}

type ListScriptsByFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsByFlowResponse) GoString() string {
	return s.String()
}

func (s *ListScriptsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptsByFlowResponse) GetBody() *ListScriptsByFlowResponseBody {
	return s.Body
}

func (s *ListScriptsByFlowResponse) SetHeaders(v map[string]*string) *ListScriptsByFlowResponse {
	s.Headers = v
	return s
}

func (s *ListScriptsByFlowResponse) SetStatusCode(v int32) *ListScriptsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptsByFlowResponse) SetBody(v *ListScriptsByFlowResponseBody) *ListScriptsByFlowResponse {
	s.Body = v
	return s
}

func (s *ListScriptsByFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
