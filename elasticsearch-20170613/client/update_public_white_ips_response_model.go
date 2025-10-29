// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicWhiteIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublicWhiteIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublicWhiteIpsResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublicWhiteIpsResponseBody) *UpdatePublicWhiteIpsResponse
	GetBody() *UpdatePublicWhiteIpsResponseBody
}

type UpdatePublicWhiteIpsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicWhiteIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublicWhiteIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublicWhiteIpsResponse) GetBody() *UpdatePublicWhiteIpsResponseBody {
	return s.Body
}

func (s *UpdatePublicWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdatePublicWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicWhiteIpsResponse) SetStatusCode(v int32) *UpdatePublicWhiteIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicWhiteIpsResponse) SetBody(v *UpdatePublicWhiteIpsResponseBody) *UpdatePublicWhiteIpsResponse {
	s.Body = v
	return s
}

func (s *UpdatePublicWhiteIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
