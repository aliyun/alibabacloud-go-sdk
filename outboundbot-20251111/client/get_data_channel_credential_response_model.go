// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataChannelCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataChannelCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetDataChannelCredentialResponseBody) *GetDataChannelCredentialResponse
	GetBody() *GetDataChannelCredentialResponseBody
}

type GetDataChannelCredentialResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataChannelCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataChannelCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataChannelCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataChannelCredentialResponse) GetBody() *GetDataChannelCredentialResponseBody {
	return s.Body
}

func (s *GetDataChannelCredentialResponse) SetHeaders(v map[string]*string) *GetDataChannelCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetDataChannelCredentialResponse) SetStatusCode(v int32) *GetDataChannelCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataChannelCredentialResponse) SetBody(v *GetDataChannelCredentialResponseBody) *GetDataChannelCredentialResponse {
	s.Body = v
	return s
}

func (s *GetDataChannelCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
