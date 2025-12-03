// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeColdStorageSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeColdStorageSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeColdStorageSizeResponse
	GetStatusCode() *int32
	SetBody(v *ResizeColdStorageSizeResponseBody) *ResizeColdStorageSizeResponse
	GetBody() *ResizeColdStorageSizeResponseBody
}

type ResizeColdStorageSizeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeColdStorageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeColdStorageSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeColdStorageSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeColdStorageSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeColdStorageSizeResponse) GetBody() *ResizeColdStorageSizeResponseBody {
	return s.Body
}

func (s *ResizeColdStorageSizeResponse) SetHeaders(v map[string]*string) *ResizeColdStorageSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeColdStorageSizeResponse) SetStatusCode(v int32) *ResizeColdStorageSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeColdStorageSizeResponse) SetBody(v *ResizeColdStorageSizeResponseBody) *ResizeColdStorageSizeResponse {
	s.Body = v
	return s
}

func (s *ResizeColdStorageSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
