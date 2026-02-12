// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerAccumulateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsConsumerAccumulateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsConsumerAccumulateResponse
	GetStatusCode() *int32
	SetBody(v *OnsConsumerAccumulateResponseBody) *OnsConsumerAccumulateResponse
	GetBody() *OnsConsumerAccumulateResponseBody
}

type OnsConsumerAccumulateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerAccumulateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerAccumulateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsConsumerAccumulateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsConsumerAccumulateResponse) GetBody() *OnsConsumerAccumulateResponseBody {
	return s.Body
}

func (s *OnsConsumerAccumulateResponse) SetHeaders(v map[string]*string) *OnsConsumerAccumulateResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerAccumulateResponse) SetStatusCode(v int32) *OnsConsumerAccumulateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerAccumulateResponse) SetBody(v *OnsConsumerAccumulateResponseBody) *OnsConsumerAccumulateResponse {
	s.Body = v
	return s
}

func (s *OnsConsumerAccumulateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
