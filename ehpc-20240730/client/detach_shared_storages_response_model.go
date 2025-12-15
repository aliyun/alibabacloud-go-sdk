// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSharedStoragesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachSharedStoragesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachSharedStoragesResponse
	GetStatusCode() *int32
	SetBody(v *DetachSharedStoragesResponseBody) *DetachSharedStoragesResponse
	GetBody() *DetachSharedStoragesResponseBody
}

type DetachSharedStoragesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachSharedStoragesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachSharedStoragesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachSharedStoragesResponse) GetBody() *DetachSharedStoragesResponseBody {
	return s.Body
}

func (s *DetachSharedStoragesResponse) SetHeaders(v map[string]*string) *DetachSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *DetachSharedStoragesResponse) SetStatusCode(v int32) *DetachSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachSharedStoragesResponse) SetBody(v *DetachSharedStoragesResponseBody) *DetachSharedStoragesResponse {
	s.Body = v
	return s
}

func (s *DetachSharedStoragesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
