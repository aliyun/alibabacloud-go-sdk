// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGwConsumerOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGwConsumerOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGwConsumerOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateGwConsumerOrderResponseBody) *CreateGwConsumerOrderResponse
	GetBody() *CreateGwConsumerOrderResponseBody
}

type CreateGwConsumerOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGwConsumerOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGwConsumerOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGwConsumerOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateGwConsumerOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGwConsumerOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGwConsumerOrderResponse) GetBody() *CreateGwConsumerOrderResponseBody {
	return s.Body
}

func (s *CreateGwConsumerOrderResponse) SetHeaders(v map[string]*string) *CreateGwConsumerOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateGwConsumerOrderResponse) SetStatusCode(v int32) *CreateGwConsumerOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGwConsumerOrderResponse) SetBody(v *CreateGwConsumerOrderResponseBody) *CreateGwConsumerOrderResponse {
	s.Body = v
	return s
}

func (s *CreateGwConsumerOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
