// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBindingCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBindingCountResponse
	GetStatusCode() *int32
	SetBody(v *GetBindingCountResponseBody) *GetBindingCountResponse
	GetBody() *GetBindingCountResponseBody
}

type GetBindingCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBindingCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBindingCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBindingCountResponse) GoString() string {
	return s.String()
}

func (s *GetBindingCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBindingCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBindingCountResponse) GetBody() *GetBindingCountResponseBody {
	return s.Body
}

func (s *GetBindingCountResponse) SetHeaders(v map[string]*string) *GetBindingCountResponse {
	s.Headers = v
	return s
}

func (s *GetBindingCountResponse) SetStatusCode(v int32) *GetBindingCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBindingCountResponse) SetBody(v *GetBindingCountResponseBody) *GetBindingCountResponse {
	s.Body = v
	return s
}

func (s *GetBindingCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
