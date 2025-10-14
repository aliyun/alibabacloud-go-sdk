// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePdnsAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePdnsAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *RemovePdnsAppKeyResponseBody) *RemovePdnsAppKeyResponse
	GetBody() *RemovePdnsAppKeyResponseBody
}

type RemovePdnsAppKeyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePdnsAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePdnsAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsAppKeyResponse) GoString() string {
	return s.String()
}

func (s *RemovePdnsAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePdnsAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePdnsAppKeyResponse) GetBody() *RemovePdnsAppKeyResponseBody {
	return s.Body
}

func (s *RemovePdnsAppKeyResponse) SetHeaders(v map[string]*string) *RemovePdnsAppKeyResponse {
	s.Headers = v
	return s
}

func (s *RemovePdnsAppKeyResponse) SetStatusCode(v int32) *RemovePdnsAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePdnsAppKeyResponse) SetBody(v *RemovePdnsAppKeyResponseBody) *RemovePdnsAppKeyResponse {
	s.Body = v
	return s
}

func (s *RemovePdnsAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
