// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpareIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSpareIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSpareIpsResponse
	GetStatusCode() *int32
	SetBody(v *ListSpareIpsResponseBody) *ListSpareIpsResponse
	GetBody() *ListSpareIpsResponseBody
}

type ListSpareIpsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSpareIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSpareIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSpareIpsResponse) GetBody() *ListSpareIpsResponseBody {
	return s.Body
}

func (s *ListSpareIpsResponse) SetHeaders(v map[string]*string) *ListSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *ListSpareIpsResponse) SetStatusCode(v int32) *ListSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpareIpsResponse) SetBody(v *ListSpareIpsResponseBody) *ListSpareIpsResponse {
	s.Body = v
	return s
}

func (s *ListSpareIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
