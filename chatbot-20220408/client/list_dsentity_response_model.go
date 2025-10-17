// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDSEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDSEntityResponse
	GetStatusCode() *int32
	SetBody(v *ListDSEntityResponseBody) *ListDSEntityResponse
	GetBody() *ListDSEntityResponseBody
}

type ListDSEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDSEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityResponse) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDSEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDSEntityResponse) GetBody() *ListDSEntityResponseBody {
	return s.Body
}

func (s *ListDSEntityResponse) SetHeaders(v map[string]*string) *ListDSEntityResponse {
	s.Headers = v
	return s
}

func (s *ListDSEntityResponse) SetStatusCode(v int32) *ListDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDSEntityResponse) SetBody(v *ListDSEntityResponseBody) *ListDSEntityResponse {
	s.Body = v
	return s
}

func (s *ListDSEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
