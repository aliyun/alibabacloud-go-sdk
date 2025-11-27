// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachRCDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachRCDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachRCDiskResponse
	GetStatusCode() *int32
	SetBody(v *DetachRCDiskResponseBody) *DetachRCDiskResponse
	GetBody() *DetachRCDiskResponseBody
}

type DetachRCDiskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachRCDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachRCDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachRCDiskResponse) GoString() string {
	return s.String()
}

func (s *DetachRCDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachRCDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachRCDiskResponse) GetBody() *DetachRCDiskResponseBody {
	return s.Body
}

func (s *DetachRCDiskResponse) SetHeaders(v map[string]*string) *DetachRCDiskResponse {
	s.Headers = v
	return s
}

func (s *DetachRCDiskResponse) SetStatusCode(v int32) *DetachRCDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachRCDiskResponse) SetBody(v *DetachRCDiskResponseBody) *DetachRCDiskResponse {
	s.Body = v
	return s
}

func (s *DetachRCDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
