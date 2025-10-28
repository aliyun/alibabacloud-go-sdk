// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *SetResourceTypeResponseBody) *SetResourceTypeResponse
	GetBody() *SetResourceTypeResponseBody
}

type SetResourceTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *SetResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetResourceTypeResponse) GetBody() *SetResourceTypeResponseBody {
	return s.Body
}

func (s *SetResourceTypeResponse) SetHeaders(v map[string]*string) *SetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *SetResourceTypeResponse) SetStatusCode(v int32) *SetResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetResourceTypeResponse) SetBody(v *SetResourceTypeResponseBody) *SetResourceTypeResponse {
	s.Body = v
	return s
}

func (s *SetResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
