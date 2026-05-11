// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySilenceTimeoutConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySilenceTimeoutConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySilenceTimeoutConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifySilenceTimeoutConfigResponseBody) *ModifySilenceTimeoutConfigResponse
	GetBody() *ModifySilenceTimeoutConfigResponseBody
}

type ModifySilenceTimeoutConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySilenceTimeoutConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySilenceTimeoutConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySilenceTimeoutConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySilenceTimeoutConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySilenceTimeoutConfigResponse) GetBody() *ModifySilenceTimeoutConfigResponseBody {
	return s.Body
}

func (s *ModifySilenceTimeoutConfigResponse) SetHeaders(v map[string]*string) *ModifySilenceTimeoutConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySilenceTimeoutConfigResponse) SetStatusCode(v int32) *ModifySilenceTimeoutConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySilenceTimeoutConfigResponse) SetBody(v *ModifySilenceTimeoutConfigResponseBody) *ModifySilenceTimeoutConfigResponse {
	s.Body = v
	return s
}

func (s *ModifySilenceTimeoutConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
