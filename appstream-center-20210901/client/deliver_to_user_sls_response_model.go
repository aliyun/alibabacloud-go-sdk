// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliverToUserSlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeliverToUserSlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeliverToUserSlsResponse
	GetStatusCode() *int32
	SetBody(v *DeliverToUserSlsResponseBody) *DeliverToUserSlsResponse
	GetBody() *DeliverToUserSlsResponseBody
}

type DeliverToUserSlsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliverToUserSlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeliverToUserSlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeliverToUserSlsResponse) GoString() string {
	return s.String()
}

func (s *DeliverToUserSlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeliverToUserSlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeliverToUserSlsResponse) GetBody() *DeliverToUserSlsResponseBody {
	return s.Body
}

func (s *DeliverToUserSlsResponse) SetHeaders(v map[string]*string) *DeliverToUserSlsResponse {
	s.Headers = v
	return s
}

func (s *DeliverToUserSlsResponse) SetStatusCode(v int32) *DeliverToUserSlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliverToUserSlsResponse) SetBody(v *DeliverToUserSlsResponseBody) *DeliverToUserSlsResponse {
	s.Body = v
	return s
}

func (s *DeliverToUserSlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
