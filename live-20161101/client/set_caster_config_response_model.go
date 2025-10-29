// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCasterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCasterConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetCasterConfigResponseBody) *SetCasterConfigResponse
	GetBody() *SetCasterConfigResponseBody
}

type SetCasterConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCasterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCasterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCasterConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCasterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCasterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCasterConfigResponse) GetBody() *SetCasterConfigResponseBody {
	return s.Body
}

func (s *SetCasterConfigResponse) SetHeaders(v map[string]*string) *SetCasterConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCasterConfigResponse) SetStatusCode(v int32) *SetCasterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCasterConfigResponse) SetBody(v *SetCasterConfigResponseBody) *SetCasterConfigResponse {
	s.Body = v
	return s
}

func (s *SetCasterConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
