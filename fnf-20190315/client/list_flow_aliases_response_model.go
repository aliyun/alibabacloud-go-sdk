// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowAliasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowAliasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowAliasesResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowAliasesResponseBody) *ListFlowAliasesResponse
	GetBody() *ListFlowAliasesResponseBody
}

type ListFlowAliasesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowAliasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListFlowAliasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowAliasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowAliasesResponse) GetBody() *ListFlowAliasesResponseBody {
	return s.Body
}

func (s *ListFlowAliasesResponse) SetHeaders(v map[string]*string) *ListFlowAliasesResponse {
	s.Headers = v
	return s
}

func (s *ListFlowAliasesResponse) SetStatusCode(v int32) *ListFlowAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowAliasesResponse) SetBody(v *ListFlowAliasesResponseBody) *ListFlowAliasesResponse {
	s.Body = v
	return s
}

func (s *ListFlowAliasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
