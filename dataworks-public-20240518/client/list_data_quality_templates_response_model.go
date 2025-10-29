// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityTemplatesResponseBody) *ListDataQualityTemplatesResponse
	GetBody() *ListDataQualityTemplatesResponseBody
}

type ListDataQualityTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityTemplatesResponse) GetBody() *ListDataQualityTemplatesResponseBody {
	return s.Body
}

func (s *ListDataQualityTemplatesResponse) SetHeaders(v map[string]*string) *ListDataQualityTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityTemplatesResponse) SetStatusCode(v int32) *ListDataQualityTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityTemplatesResponse) SetBody(v *ListDataQualityTemplatesResponseBody) *ListDataQualityTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
