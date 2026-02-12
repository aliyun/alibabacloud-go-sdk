// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsConsumerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsConsumerStatusResponse
	GetStatusCode() *int32
	SetBody(v *OnsConsumerStatusResponseBody) *OnsConsumerStatusResponse
	GetBody() *OnsConsumerStatusResponseBody
}

type OnsConsumerStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsConsumerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsConsumerStatusResponse) GetBody() *OnsConsumerStatusResponseBody {
	return s.Body
}

func (s *OnsConsumerStatusResponse) SetHeaders(v map[string]*string) *OnsConsumerStatusResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerStatusResponse) SetStatusCode(v int32) *OnsConsumerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerStatusResponse) SetBody(v *OnsConsumerStatusResponseBody) *OnsConsumerStatusResponse {
	s.Body = v
	return s
}

func (s *OnsConsumerStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
