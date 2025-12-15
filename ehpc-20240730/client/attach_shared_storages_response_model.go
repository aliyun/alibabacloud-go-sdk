// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSharedStoragesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachSharedStoragesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachSharedStoragesResponse
	GetStatusCode() *int32
	SetBody(v *AttachSharedStoragesResponseBody) *AttachSharedStoragesResponse
	GetBody() *AttachSharedStoragesResponseBody
}

type AttachSharedStoragesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachSharedStoragesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachSharedStoragesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachSharedStoragesResponse) GetBody() *AttachSharedStoragesResponseBody {
	return s.Body
}

func (s *AttachSharedStoragesResponse) SetHeaders(v map[string]*string) *AttachSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *AttachSharedStoragesResponse) SetStatusCode(v int32) *AttachSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachSharedStoragesResponse) SetBody(v *AttachSharedStoragesResponseBody) *AttachSharedStoragesResponse {
	s.Body = v
	return s
}

func (s *AttachSharedStoragesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
