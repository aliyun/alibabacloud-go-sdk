// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByAILabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaByAILabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaByAILabelResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaByAILabelResponseBody) *SearchMediaByAILabelResponse
	GetBody() *SearchMediaByAILabelResponseBody
}

type SearchMediaByAILabelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaByAILabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaByAILabelResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaByAILabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaByAILabelResponse) GetBody() *SearchMediaByAILabelResponseBody {
	return s.Body
}

func (s *SearchMediaByAILabelResponse) SetHeaders(v map[string]*string) *SearchMediaByAILabelResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaByAILabelResponse) SetStatusCode(v int32) *SearchMediaByAILabelResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaByAILabelResponse) SetBody(v *SearchMediaByAILabelResponseBody) *SearchMediaByAILabelResponse {
	s.Body = v
	return s
}

func (s *SearchMediaByAILabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
