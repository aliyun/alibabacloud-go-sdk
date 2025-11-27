// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualRetrievalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContextualRetrievalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContextualRetrievalResponse
	GetStatusCode() *int32
	SetBody(v *ContextualRetrievalResponseBody) *ContextualRetrievalResponse
	GetBody() *ContextualRetrievalResponseBody
}

type ContextualRetrievalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContextualRetrievalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContextualRetrievalResponse) String() string {
	return dara.Prettify(s)
}

func (s ContextualRetrievalResponse) GoString() string {
	return s.String()
}

func (s *ContextualRetrievalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContextualRetrievalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContextualRetrievalResponse) GetBody() *ContextualRetrievalResponseBody {
	return s.Body
}

func (s *ContextualRetrievalResponse) SetHeaders(v map[string]*string) *ContextualRetrievalResponse {
	s.Headers = v
	return s
}

func (s *ContextualRetrievalResponse) SetStatusCode(v int32) *ContextualRetrievalResponse {
	s.StatusCode = &v
	return s
}

func (s *ContextualRetrievalResponse) SetBody(v *ContextualRetrievalResponseBody) *ContextualRetrievalResponse {
	s.Body = v
	return s
}

func (s *ContextualRetrievalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
