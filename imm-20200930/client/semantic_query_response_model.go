// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSemanticQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SemanticQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SemanticQueryResponse
	GetStatusCode() *int32
	SetBody(v *SemanticQueryResponseBody) *SemanticQueryResponse
	GetBody() *SemanticQueryResponseBody
}

type SemanticQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SemanticQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SemanticQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s SemanticQueryResponse) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SemanticQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SemanticQueryResponse) GetBody() *SemanticQueryResponseBody {
	return s.Body
}

func (s *SemanticQueryResponse) SetHeaders(v map[string]*string) *SemanticQueryResponse {
	s.Headers = v
	return s
}

func (s *SemanticQueryResponse) SetStatusCode(v int32) *SemanticQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SemanticQueryResponse) SetBody(v *SemanticQueryResponseBody) *SemanticQueryResponse {
	s.Body = v
	return s
}

func (s *SemanticQueryResponse) Validate() error {
	return dara.Validate(s)
}
