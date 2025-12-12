// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityCheckSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityCheckSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityCheckSchemeResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityCheckSchemeResponseBody) *ListQualityCheckSchemeResponse
	GetBody() *ListQualityCheckSchemeResponseBody
}

type ListQualityCheckSchemeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityCheckSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityCheckSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityCheckSchemeResponse) GetBody() *ListQualityCheckSchemeResponseBody {
	return s.Body
}

func (s *ListQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *ListQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetStatusCode(v int32) *ListQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetBody(v *ListQualityCheckSchemeResponseBody) *ListQualityCheckSchemeResponse {
	s.Body = v
	return s
}

func (s *ListQualityCheckSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
