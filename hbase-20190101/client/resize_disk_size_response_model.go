// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeDiskSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeDiskSizeResponse
	GetStatusCode() *int32
	SetBody(v *ResizeDiskSizeResponseBody) *ResizeDiskSizeResponse
	GetBody() *ResizeDiskSizeResponseBody
}

type ResizeDiskSizeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeDiskSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeDiskSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeDiskSizeResponse) GetBody() *ResizeDiskSizeResponseBody {
	return s.Body
}

func (s *ResizeDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeDiskSizeResponse) SetStatusCode(v int32) *ResizeDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeDiskSizeResponse) SetBody(v *ResizeDiskSizeResponseBody) *ResizeDiskSizeResponse {
	s.Body = v
	return s
}

func (s *ResizeDiskSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
