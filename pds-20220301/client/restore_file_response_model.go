// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreFileResponse
	GetStatusCode() *int32
	SetBody(v *RestoreFileResponseBody) *RestoreFileResponse
	GetBody() *RestoreFileResponseBody
}

type RestoreFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreFileResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreFileResponse) GoString() string {
	return s.String()
}

func (s *RestoreFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreFileResponse) GetBody() *RestoreFileResponseBody {
	return s.Body
}

func (s *RestoreFileResponse) SetHeaders(v map[string]*string) *RestoreFileResponse {
	s.Headers = v
	return s
}

func (s *RestoreFileResponse) SetStatusCode(v int32) *RestoreFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreFileResponse) SetBody(v *RestoreFileResponseBody) *RestoreFileResponse {
	s.Body = v
	return s
}

func (s *RestoreFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
