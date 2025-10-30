// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxbResponse
	GetStatusCode() *int32
	SetBody(v *BindAxbResponseBody) *BindAxbResponse
	GetBody() *BindAxbResponseBody
}

type BindAxbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxbResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxbResponse) GoString() string {
	return s.String()
}

func (s *BindAxbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxbResponse) GetBody() *BindAxbResponseBody {
	return s.Body
}

func (s *BindAxbResponse) SetHeaders(v map[string]*string) *BindAxbResponse {
	s.Headers = v
	return s
}

func (s *BindAxbResponse) SetStatusCode(v int32) *BindAxbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxbResponse) SetBody(v *BindAxbResponseBody) *BindAxbResponse {
	s.Body = v
	return s
}

func (s *BindAxbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
