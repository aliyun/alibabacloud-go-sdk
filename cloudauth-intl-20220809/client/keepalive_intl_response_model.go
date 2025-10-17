// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeepaliveIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KeepaliveIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KeepaliveIntlResponse
	GetStatusCode() *int32
	SetBody(v *KeepaliveIntlResponseBody) *KeepaliveIntlResponse
	GetBody() *KeepaliveIntlResponseBody
}

type KeepaliveIntlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KeepaliveIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KeepaliveIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s KeepaliveIntlResponse) GoString() string {
	return s.String()
}

func (s *KeepaliveIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KeepaliveIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KeepaliveIntlResponse) GetBody() *KeepaliveIntlResponseBody {
	return s.Body
}

func (s *KeepaliveIntlResponse) SetHeaders(v map[string]*string) *KeepaliveIntlResponse {
	s.Headers = v
	return s
}

func (s *KeepaliveIntlResponse) SetStatusCode(v int32) *KeepaliveIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *KeepaliveIntlResponse) SetBody(v *KeepaliveIntlResponseBody) *KeepaliveIntlResponse {
	s.Body = v
	return s
}

func (s *KeepaliveIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
