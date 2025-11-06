// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthConsumerDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayAuthConsumerDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayAuthConsumerDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayAuthConsumerDetailResponseBody) *GetGatewayAuthConsumerDetailResponse
	GetBody() *GetGatewayAuthConsumerDetailResponseBody
}

type GetGatewayAuthConsumerDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayAuthConsumerDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayAuthConsumerDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthConsumerDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthConsumerDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayAuthConsumerDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayAuthConsumerDetailResponse) GetBody() *GetGatewayAuthConsumerDetailResponseBody {
	return s.Body
}

func (s *GetGatewayAuthConsumerDetailResponse) SetHeaders(v map[string]*string) *GetGatewayAuthConsumerDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponse) SetStatusCode(v int32) *GetGatewayAuthConsumerDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponse) SetBody(v *GetGatewayAuthConsumerDetailResponseBody) *GetGatewayAuthConsumerDetailResponse {
	s.Body = v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
