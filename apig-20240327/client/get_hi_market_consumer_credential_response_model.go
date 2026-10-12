// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHiMarketConsumerCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHiMarketConsumerCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetHiMarketConsumerCredentialResponseBody) *GetHiMarketConsumerCredentialResponse
	GetBody() *GetHiMarketConsumerCredentialResponseBody
}

type GetHiMarketConsumerCredentialResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHiMarketConsumerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHiMarketConsumerCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHiMarketConsumerCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHiMarketConsumerCredentialResponse) GetBody() *GetHiMarketConsumerCredentialResponseBody {
	return s.Body
}

func (s *GetHiMarketConsumerCredentialResponse) SetHeaders(v map[string]*string) *GetHiMarketConsumerCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetHiMarketConsumerCredentialResponse) SetStatusCode(v int32) *GetHiMarketConsumerCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHiMarketConsumerCredentialResponse) SetBody(v *GetHiMarketConsumerCredentialResponseBody) *GetHiMarketConsumerCredentialResponse {
	s.Body = v
	return s
}

func (s *GetHiMarketConsumerCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
