// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConvertableEcuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConvertableEcuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConvertableEcuResponse
	GetStatusCode() *int32
	SetBody(v *ListConvertableEcuResponseBody) *ListConvertableEcuResponse
	GetBody() *ListConvertableEcuResponseBody
}

type ListConvertableEcuResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConvertableEcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConvertableEcuResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConvertableEcuResponse) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConvertableEcuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConvertableEcuResponse) GetBody() *ListConvertableEcuResponseBody {
	return s.Body
}

func (s *ListConvertableEcuResponse) SetHeaders(v map[string]*string) *ListConvertableEcuResponse {
	s.Headers = v
	return s
}

func (s *ListConvertableEcuResponse) SetStatusCode(v int32) *ListConvertableEcuResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConvertableEcuResponse) SetBody(v *ListConvertableEcuResponseBody) *ListConvertableEcuResponse {
	s.Body = v
	return s
}

func (s *ListConvertableEcuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
