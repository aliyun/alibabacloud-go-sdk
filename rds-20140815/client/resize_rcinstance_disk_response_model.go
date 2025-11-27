// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeRCInstanceDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeRCInstanceDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeRCInstanceDiskResponse
	GetStatusCode() *int32
	SetBody(v *ResizeRCInstanceDiskResponseBody) *ResizeRCInstanceDiskResponse
	GetBody() *ResizeRCInstanceDiskResponseBody
}

type ResizeRCInstanceDiskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeRCInstanceDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeRCInstanceDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeRCInstanceDiskResponse) GoString() string {
	return s.String()
}

func (s *ResizeRCInstanceDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeRCInstanceDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeRCInstanceDiskResponse) GetBody() *ResizeRCInstanceDiskResponseBody {
	return s.Body
}

func (s *ResizeRCInstanceDiskResponse) SetHeaders(v map[string]*string) *ResizeRCInstanceDiskResponse {
	s.Headers = v
	return s
}

func (s *ResizeRCInstanceDiskResponse) SetStatusCode(v int32) *ResizeRCInstanceDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeRCInstanceDiskResponse) SetBody(v *ResizeRCInstanceDiskResponseBody) *ResizeRCInstanceDiskResponse {
	s.Body = v
	return s
}

func (s *ResizeRCInstanceDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
