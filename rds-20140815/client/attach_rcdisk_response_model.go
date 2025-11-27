// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachRCDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachRCDiskResponse
	GetStatusCode() *int32
	SetBody(v *AttachRCDiskResponseBody) *AttachRCDiskResponse
	GetBody() *AttachRCDiskResponseBody
}

type AttachRCDiskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachRCDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachRCDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachRCDiskResponse) GoString() string {
	return s.String()
}

func (s *AttachRCDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachRCDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachRCDiskResponse) GetBody() *AttachRCDiskResponseBody {
	return s.Body
}

func (s *AttachRCDiskResponse) SetHeaders(v map[string]*string) *AttachRCDiskResponse {
	s.Headers = v
	return s
}

func (s *AttachRCDiskResponse) SetStatusCode(v int32) *AttachRCDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachRCDiskResponse) SetBody(v *AttachRCDiskResponseBody) *AttachRCDiskResponse {
	s.Body = v
	return s
}

func (s *AttachRCDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
