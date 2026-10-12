// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHiMarketConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHiMarketConsumerResponse
	GetStatusCode() *int32
	SetBody(v *GetHiMarketConsumerResponseBody) *GetHiMarketConsumerResponse
	GetBody() *GetHiMarketConsumerResponseBody
}

type GetHiMarketConsumerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHiMarketConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHiMarketConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerResponse) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHiMarketConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHiMarketConsumerResponse) GetBody() *GetHiMarketConsumerResponseBody {
	return s.Body
}

func (s *GetHiMarketConsumerResponse) SetHeaders(v map[string]*string) *GetHiMarketConsumerResponse {
	s.Headers = v
	return s
}

func (s *GetHiMarketConsumerResponse) SetStatusCode(v int32) *GetHiMarketConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHiMarketConsumerResponse) SetBody(v *GetHiMarketConsumerResponseBody) *GetHiMarketConsumerResponse {
	s.Body = v
	return s
}

func (s *GetHiMarketConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
