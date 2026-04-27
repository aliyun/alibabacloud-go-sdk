// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudInterruptIvrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudInterruptIvrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudInterruptIvrResponse
	GetStatusCode() *int32
	SetBody(v *CloudInterruptIvrResponseBody) *CloudInterruptIvrResponse
	GetBody() *CloudInterruptIvrResponseBody
}

type CloudInterruptIvrResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudInterruptIvrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudInterruptIvrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudInterruptIvrResponse) GoString() string {
	return s.String()
}

func (s *CloudInterruptIvrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudInterruptIvrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudInterruptIvrResponse) GetBody() *CloudInterruptIvrResponseBody {
	return s.Body
}

func (s *CloudInterruptIvrResponse) SetHeaders(v map[string]*string) *CloudInterruptIvrResponse {
	s.Headers = v
	return s
}

func (s *CloudInterruptIvrResponse) SetStatusCode(v int32) *CloudInterruptIvrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudInterruptIvrResponse) SetBody(v *CloudInterruptIvrResponseBody) *CloudInterruptIvrResponse {
	s.Body = v
	return s
}

func (s *CloudInterruptIvrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
