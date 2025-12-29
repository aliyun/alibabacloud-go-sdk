// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateWelcomeTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrUpdateWelcomeTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponse
	GetStatusCode() *int32
	SetBody(v *AddOrUpdateWelcomeTextResponseBody) *AddOrUpdateWelcomeTextResponse
	GetBody() *AddOrUpdateWelcomeTextResponseBody
}

type AddOrUpdateWelcomeTextResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateWelcomeTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateWelcomeTextResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateWelcomeTextResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrUpdateWelcomeTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateWelcomeTextResponse) GetBody() *AddOrUpdateWelcomeTextResponseBody {
	return s.Body
}

func (s *AddOrUpdateWelcomeTextResponse) SetHeaders(v map[string]*string) *AddOrUpdateWelcomeTextResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateWelcomeTextResponse) SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponse) SetBody(v *AddOrUpdateWelcomeTextResponseBody) *AddOrUpdateWelcomeTextResponse {
	s.Body = v
	return s
}

func (s *AddOrUpdateWelcomeTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
