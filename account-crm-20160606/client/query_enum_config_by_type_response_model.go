// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnumConfigByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEnumConfigByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEnumConfigByTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryEnumConfigByTypeResponseBody) *QueryEnumConfigByTypeResponse
	GetBody() *QueryEnumConfigByTypeResponseBody
}

type QueryEnumConfigByTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnumConfigByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnumConfigByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEnumConfigByTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryEnumConfigByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEnumConfigByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEnumConfigByTypeResponse) GetBody() *QueryEnumConfigByTypeResponseBody {
	return s.Body
}

func (s *QueryEnumConfigByTypeResponse) SetHeaders(v map[string]*string) *QueryEnumConfigByTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryEnumConfigByTypeResponse) SetStatusCode(v int32) *QueryEnumConfigByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnumConfigByTypeResponse) SetBody(v *QueryEnumConfigByTypeResponseBody) *QueryEnumConfigByTypeResponse {
	s.Body = v
	return s
}

func (s *QueryEnumConfigByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
