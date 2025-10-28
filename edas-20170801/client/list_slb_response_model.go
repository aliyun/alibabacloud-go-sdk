// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSlbResponse
	GetStatusCode() *int32
	SetBody(v *ListSlbResponseBody) *ListSlbResponse
	GetBody() *ListSlbResponseBody
}

type ListSlbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSlbResponse) GoString() string {
	return s.String()
}

func (s *ListSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSlbResponse) GetBody() *ListSlbResponseBody {
	return s.Body
}

func (s *ListSlbResponse) SetHeaders(v map[string]*string) *ListSlbResponse {
	s.Headers = v
	return s
}

func (s *ListSlbResponse) SetStatusCode(v int32) *ListSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlbResponse) SetBody(v *ListSlbResponseBody) *ListSlbResponse {
	s.Body = v
	return s
}

func (s *ListSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
