// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyFileResponse
	GetStatusCode() *int32
	SetBody(v *CopyFileResponseBody) *CopyFileResponse
	GetBody() *CopyFileResponseBody
}

type CopyFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyFileResponse) GoString() string {
	return s.String()
}

func (s *CopyFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyFileResponse) GetBody() *CopyFileResponseBody {
	return s.Body
}

func (s *CopyFileResponse) SetHeaders(v map[string]*string) *CopyFileResponse {
	s.Headers = v
	return s
}

func (s *CopyFileResponse) SetStatusCode(v int32) *CopyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyFileResponse) SetBody(v *CopyFileResponseBody) *CopyFileResponse {
	s.Body = v
	return s
}

func (s *CopyFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
