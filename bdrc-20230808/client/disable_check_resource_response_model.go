// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCheckResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCheckResourceResponse
	GetStatusCode() *int32
	SetBody(v *DisableCheckResourceResponseBody) *DisableCheckResourceResponse
	GetBody() *DisableCheckResourceResponseBody
}

type DisableCheckResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCheckResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCheckResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCheckResourceResponse) GetBody() *DisableCheckResourceResponseBody {
	return s.Body
}

func (s *DisableCheckResourceResponse) SetHeaders(v map[string]*string) *DisableCheckResourceResponse {
	s.Headers = v
	return s
}

func (s *DisableCheckResourceResponse) SetStatusCode(v int32) *DisableCheckResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCheckResourceResponse) SetBody(v *DisableCheckResourceResponseBody) *DisableCheckResourceResponse {
	s.Body = v
	return s
}

func (s *DisableCheckResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
