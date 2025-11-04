// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFileUploadLeaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyFileUploadLeaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyFileUploadLeaseResponse
	GetStatusCode() *int32
	SetBody(v *ApplyFileUploadLeaseResponseBody) *ApplyFileUploadLeaseResponse
	GetBody() *ApplyFileUploadLeaseResponseBody
}

type ApplyFileUploadLeaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFileUploadLeaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFileUploadLeaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseResponse) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyFileUploadLeaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyFileUploadLeaseResponse) GetBody() *ApplyFileUploadLeaseResponseBody {
	return s.Body
}

func (s *ApplyFileUploadLeaseResponse) SetHeaders(v map[string]*string) *ApplyFileUploadLeaseResponse {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetStatusCode(v int32) *ApplyFileUploadLeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetBody(v *ApplyFileUploadLeaseResponseBody) *ApplyFileUploadLeaseResponse {
	s.Body = v
	return s
}

func (s *ApplyFileUploadLeaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
