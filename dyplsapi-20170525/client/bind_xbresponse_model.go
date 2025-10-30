// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindXBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindXBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindXBResponse
	GetStatusCode() *int32
	SetBody(v *BindXBResponseBody) *BindXBResponse
	GetBody() *BindXBResponseBody
}

type BindXBResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindXBResponse) String() string {
	return dara.Prettify(s)
}

func (s BindXBResponse) GoString() string {
	return s.String()
}

func (s *BindXBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindXBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindXBResponse) GetBody() *BindXBResponseBody {
	return s.Body
}

func (s *BindXBResponse) SetHeaders(v map[string]*string) *BindXBResponse {
	s.Headers = v
	return s
}

func (s *BindXBResponse) SetStatusCode(v int32) *BindXBResponse {
	s.StatusCode = &v
	return s
}

func (s *BindXBResponse) SetBody(v *BindXBResponseBody) *BindXBResponse {
	s.Body = v
	return s
}

func (s *BindXBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
