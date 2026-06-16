// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryAttributeMatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CategoryAttributeMatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CategoryAttributeMatchResponse
	GetStatusCode() *int32
	SetBody(v *CategoryAttributeMatchResponseBody) *CategoryAttributeMatchResponse
	GetBody() *CategoryAttributeMatchResponseBody
}

type CategoryAttributeMatchResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoryAttributeMatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CategoryAttributeMatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchResponse) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CategoryAttributeMatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CategoryAttributeMatchResponse) GetBody() *CategoryAttributeMatchResponseBody {
	return s.Body
}

func (s *CategoryAttributeMatchResponse) SetHeaders(v map[string]*string) *CategoryAttributeMatchResponse {
	s.Headers = v
	return s
}

func (s *CategoryAttributeMatchResponse) SetStatusCode(v int32) *CategoryAttributeMatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CategoryAttributeMatchResponse) SetBody(v *CategoryAttributeMatchResponseBody) *CategoryAttributeMatchResponse {
	s.Body = v
	return s
}

func (s *CategoryAttributeMatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
