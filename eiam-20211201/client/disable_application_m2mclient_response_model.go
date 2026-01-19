// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationM2MClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationM2MClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationM2MClientResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationM2MClientResponseBody) *DisableApplicationM2MClientResponse
	GetBody() *DisableApplicationM2MClientResponseBody
}

type DisableApplicationM2MClientResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationM2MClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationM2MClientResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationM2MClientResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationM2MClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationM2MClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationM2MClientResponse) GetBody() *DisableApplicationM2MClientResponseBody {
	return s.Body
}

func (s *DisableApplicationM2MClientResponse) SetHeaders(v map[string]*string) *DisableApplicationM2MClientResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationM2MClientResponse) SetStatusCode(v int32) *DisableApplicationM2MClientResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationM2MClientResponse) SetBody(v *DisableApplicationM2MClientResponseBody) *DisableApplicationM2MClientResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationM2MClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
