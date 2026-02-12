// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerGetConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsConsumerGetConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsConsumerGetConnectionResponse
	GetStatusCode() *int32
	SetBody(v *OnsConsumerGetConnectionResponseBody) *OnsConsumerGetConnectionResponse
	GetBody() *OnsConsumerGetConnectionResponseBody
}

type OnsConsumerGetConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerGetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerGetConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsConsumerGetConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsConsumerGetConnectionResponse) GetBody() *OnsConsumerGetConnectionResponseBody {
	return s.Body
}

func (s *OnsConsumerGetConnectionResponse) SetHeaders(v map[string]*string) *OnsConsumerGetConnectionResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerGetConnectionResponse) SetStatusCode(v int32) *OnsConsumerGetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerGetConnectionResponse) SetBody(v *OnsConsumerGetConnectionResponseBody) *OnsConsumerGetConnectionResponse {
	s.Body = v
	return s
}

func (s *OnsConsumerGetConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
