// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEcsSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindEcsSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindEcsSlbResponse
	GetStatusCode() *int32
	SetBody(v *BindEcsSlbResponseBody) *BindEcsSlbResponse
	GetBody() *BindEcsSlbResponseBody
}

type BindEcsSlbResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEcsSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEcsSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s BindEcsSlbResponse) GoString() string {
	return s.String()
}

func (s *BindEcsSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindEcsSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindEcsSlbResponse) GetBody() *BindEcsSlbResponseBody {
	return s.Body
}

func (s *BindEcsSlbResponse) SetHeaders(v map[string]*string) *BindEcsSlbResponse {
	s.Headers = v
	return s
}

func (s *BindEcsSlbResponse) SetStatusCode(v int32) *BindEcsSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEcsSlbResponse) SetBody(v *BindEcsSlbResponseBody) *BindEcsSlbResponse {
	s.Body = v
	return s
}

func (s *BindEcsSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
