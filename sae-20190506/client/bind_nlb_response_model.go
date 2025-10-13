// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindNlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindNlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindNlbResponse
	GetStatusCode() *int32
	SetBody(v *BindNlbResponseBody) *BindNlbResponse
	GetBody() *BindNlbResponseBody
}

type BindNlbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindNlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindNlbResponse) String() string {
	return dara.Prettify(s)
}

func (s BindNlbResponse) GoString() string {
	return s.String()
}

func (s *BindNlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindNlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindNlbResponse) GetBody() *BindNlbResponseBody {
	return s.Body
}

func (s *BindNlbResponse) SetHeaders(v map[string]*string) *BindNlbResponse {
	s.Headers = v
	return s
}

func (s *BindNlbResponse) SetStatusCode(v int32) *BindNlbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindNlbResponse) SetBody(v *BindNlbResponseBody) *BindNlbResponse {
	s.Body = v
	return s
}

func (s *BindNlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
