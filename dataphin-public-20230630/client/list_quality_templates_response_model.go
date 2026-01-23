// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityTemplatesResponseBody) *ListQualityTemplatesResponse
	GetBody() *ListQualityTemplatesResponseBody
}

type ListQualityTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityTemplatesResponse) GetBody() *ListQualityTemplatesResponseBody {
	return s.Body
}

func (s *ListQualityTemplatesResponse) SetHeaders(v map[string]*string) *ListQualityTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListQualityTemplatesResponse) SetStatusCode(v int32) *ListQualityTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityTemplatesResponse) SetBody(v *ListQualityTemplatesResponseBody) *ListQualityTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListQualityTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
