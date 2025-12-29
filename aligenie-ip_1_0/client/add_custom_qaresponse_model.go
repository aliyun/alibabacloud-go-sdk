// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomQAResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomQAResponseBody) *AddCustomQAResponse
	GetBody() *AddCustomQAResponseBody
}

type AddCustomQAResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomQAResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAResponse) GoString() string {
	return s.String()
}

func (s *AddCustomQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomQAResponse) GetBody() *AddCustomQAResponseBody {
	return s.Body
}

func (s *AddCustomQAResponse) SetHeaders(v map[string]*string) *AddCustomQAResponse {
	s.Headers = v
	return s
}

func (s *AddCustomQAResponse) SetStatusCode(v int32) *AddCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAResponse) SetBody(v *AddCustomQAResponseBody) *AddCustomQAResponse {
	s.Body = v
	return s
}

func (s *AddCustomQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
