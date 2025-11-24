// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKubeconfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeKubeconfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeKubeconfigResponse
	GetStatusCode() *int32
	SetBody(v *RevokeKubeconfigResponseBody) *RevokeKubeconfigResponse
	GetBody() *RevokeKubeconfigResponseBody
}

type RevokeKubeconfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeKubeconfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeKubeconfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeKubeconfigResponse) GetBody() *RevokeKubeconfigResponseBody {
	return s.Body
}

func (s *RevokeKubeconfigResponse) SetHeaders(v map[string]*string) *RevokeKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *RevokeKubeconfigResponse) SetStatusCode(v int32) *RevokeKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeKubeconfigResponse) SetBody(v *RevokeKubeconfigResponseBody) *RevokeKubeconfigResponse {
	s.Body = v
	return s
}

func (s *RevokeKubeconfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
