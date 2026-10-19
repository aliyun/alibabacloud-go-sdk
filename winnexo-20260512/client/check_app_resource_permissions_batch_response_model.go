// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAppResourcePermissionsBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAppResourcePermissionsBatchResponse
	GetStatusCode() *int32
	SetBody(v *CheckAppResourcePermissionsBatchResponseBody) *CheckAppResourcePermissionsBatchResponse
	GetBody() *CheckAppResourcePermissionsBatchResponseBody
}

type CheckAppResourcePermissionsBatchResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAppResourcePermissionsBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAppResourcePermissionsBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchResponse) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAppResourcePermissionsBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAppResourcePermissionsBatchResponse) GetBody() *CheckAppResourcePermissionsBatchResponseBody {
	return s.Body
}

func (s *CheckAppResourcePermissionsBatchResponse) SetHeaders(v map[string]*string) *CheckAppResourcePermissionsBatchResponse {
	s.Headers = v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponse) SetStatusCode(v int32) *CheckAppResourcePermissionsBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponse) SetBody(v *CheckAppResourcePermissionsBatchResponseBody) *CheckAppResourcePermissionsBatchResponse {
	s.Body = v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
