// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSerialNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSerialNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSerialNumberResponse
	GetStatusCode() *int32
	SetBody(v *BindSerialNumberResponseBody) *BindSerialNumberResponse
	GetBody() *BindSerialNumberResponseBody
}

type BindSerialNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSerialNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *BindSerialNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSerialNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSerialNumberResponse) GetBody() *BindSerialNumberResponseBody {
	return s.Body
}

func (s *BindSerialNumberResponse) SetHeaders(v map[string]*string) *BindSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *BindSerialNumberResponse) SetStatusCode(v int32) *BindSerialNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSerialNumberResponse) SetBody(v *BindSerialNumberResponseBody) *BindSerialNumberResponse {
	s.Body = v
	return s
}

func (s *BindSerialNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
