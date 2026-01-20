// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSharingWithResourceDirectoryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSharingWithResourceDirectoryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSharingWithResourceDirectoryStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckSharingWithResourceDirectoryStatusResponseBody) *CheckSharingWithResourceDirectoryStatusResponse
	GetBody() *CheckSharingWithResourceDirectoryStatusResponseBody
}

type CheckSharingWithResourceDirectoryStatusResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSharingWithResourceDirectoryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSharingWithResourceDirectoryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSharingWithResourceDirectoryStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) GetBody() *CheckSharingWithResourceDirectoryStatusResponseBody {
	return s.Body
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetHeaders(v map[string]*string) *CheckSharingWithResourceDirectoryStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetStatusCode(v int32) *CheckSharingWithResourceDirectoryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetBody(v *CheckSharingWithResourceDirectoryStatusResponseBody) *CheckSharingWithResourceDirectoryStatusResponse {
	s.Body = v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
