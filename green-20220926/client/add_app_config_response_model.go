// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddAppConfigResponseBody) *AddAppConfigResponse
	GetBody() *AddAppConfigResponseBody
}

type AddAppConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAppConfigResponse) GoString() string {
	return s.String()
}

func (s *AddAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAppConfigResponse) GetBody() *AddAppConfigResponseBody {
	return s.Body
}

func (s *AddAppConfigResponse) SetHeaders(v map[string]*string) *AddAppConfigResponse {
	s.Headers = v
	return s
}

func (s *AddAppConfigResponse) SetStatusCode(v int32) *AddAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAppConfigResponse) SetBody(v *AddAppConfigResponseBody) *AddAppConfigResponse {
	s.Body = v
	return s
}

func (s *AddAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
