// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityScansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityScansResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityScansResponseBody) *ListDataQualityScansResponse
	GetBody() *ListDataQualityScansResponseBody
}

type ListDataQualityScansResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityScansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityScansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityScansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityScansResponse) GetBody() *ListDataQualityScansResponseBody {
	return s.Body
}

func (s *ListDataQualityScansResponse) SetHeaders(v map[string]*string) *ListDataQualityScansResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityScansResponse) SetStatusCode(v int32) *ListDataQualityScansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityScansResponse) SetBody(v *ListDataQualityScansResponseBody) *ListDataQualityScansResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityScansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
