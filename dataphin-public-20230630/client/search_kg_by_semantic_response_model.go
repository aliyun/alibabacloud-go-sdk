// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKgBySemanticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchKgBySemanticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchKgBySemanticResponse
	GetStatusCode() *int32
	SetBody(v *SearchKgBySemanticResponseBody) *SearchKgBySemanticResponse
	GetBody() *SearchKgBySemanticResponseBody
}

type SearchKgBySemanticResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchKgBySemanticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchKgBySemanticResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticResponse) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchKgBySemanticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchKgBySemanticResponse) GetBody() *SearchKgBySemanticResponseBody {
	return s.Body
}

func (s *SearchKgBySemanticResponse) SetHeaders(v map[string]*string) *SearchKgBySemanticResponse {
	s.Headers = v
	return s
}

func (s *SearchKgBySemanticResponse) SetStatusCode(v int32) *SearchKgBySemanticResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchKgBySemanticResponse) SetBody(v *SearchKgBySemanticResponseBody) *SearchKgBySemanticResponse {
	s.Body = v
	return s
}

func (s *SearchKgBySemanticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
