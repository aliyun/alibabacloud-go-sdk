// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataMaskingColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataMaskingColumnsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataMaskingColumnsResponseBody) *ListDataMaskingColumnsResponse
	GetBody() *ListDataMaskingColumnsResponseBody
}

type ListDataMaskingColumnsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataMaskingColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataMaskingColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListDataMaskingColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataMaskingColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataMaskingColumnsResponse) GetBody() *ListDataMaskingColumnsResponseBody {
	return s.Body
}

func (s *ListDataMaskingColumnsResponse) SetHeaders(v map[string]*string) *ListDataMaskingColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListDataMaskingColumnsResponse) SetStatusCode(v int32) *ListDataMaskingColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataMaskingColumnsResponse) SetBody(v *ListDataMaskingColumnsResponseBody) *ListDataMaskingColumnsResponse {
	s.Body = v
	return s
}

func (s *ListDataMaskingColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
