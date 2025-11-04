// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTempStorageLeaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyTempStorageLeaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyTempStorageLeaseResponse
	GetStatusCode() *int32
	SetBody(v *ApplyTempStorageLeaseResponseBody) *ApplyTempStorageLeaseResponse
	GetBody() *ApplyTempStorageLeaseResponseBody
}

type ApplyTempStorageLeaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyTempStorageLeaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyTempStorageLeaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyTempStorageLeaseResponse) GoString() string {
	return s.String()
}

func (s *ApplyTempStorageLeaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyTempStorageLeaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyTempStorageLeaseResponse) GetBody() *ApplyTempStorageLeaseResponseBody {
	return s.Body
}

func (s *ApplyTempStorageLeaseResponse) SetHeaders(v map[string]*string) *ApplyTempStorageLeaseResponse {
	s.Headers = v
	return s
}

func (s *ApplyTempStorageLeaseResponse) SetStatusCode(v int32) *ApplyTempStorageLeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyTempStorageLeaseResponse) SetBody(v *ApplyTempStorageLeaseResponseBody) *ApplyTempStorageLeaseResponse {
	s.Body = v
	return s
}

func (s *ApplyTempStorageLeaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
