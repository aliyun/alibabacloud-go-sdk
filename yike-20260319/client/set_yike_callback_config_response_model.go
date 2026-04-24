// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeCallbackConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetYikeCallbackConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetYikeCallbackConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetYikeCallbackConfigResponseBody) *SetYikeCallbackConfigResponse
	GetBody() *SetYikeCallbackConfigResponseBody
}

type SetYikeCallbackConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetYikeCallbackConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetYikeCallbackConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetYikeCallbackConfigResponse) GoString() string {
	return s.String()
}

func (s *SetYikeCallbackConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetYikeCallbackConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetYikeCallbackConfigResponse) GetBody() *SetYikeCallbackConfigResponseBody {
	return s.Body
}

func (s *SetYikeCallbackConfigResponse) SetHeaders(v map[string]*string) *SetYikeCallbackConfigResponse {
	s.Headers = v
	return s
}

func (s *SetYikeCallbackConfigResponse) SetStatusCode(v int32) *SetYikeCallbackConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetYikeCallbackConfigResponse) SetBody(v *SetYikeCallbackConfigResponseBody) *SetYikeCallbackConfigResponse {
	s.Body = v
	return s
}

func (s *SetYikeCallbackConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
