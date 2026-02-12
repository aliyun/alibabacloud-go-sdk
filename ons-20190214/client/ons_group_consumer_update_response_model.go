// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupConsumerUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsGroupConsumerUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsGroupConsumerUpdateResponse
	GetStatusCode() *int32
	SetBody(v *OnsGroupConsumerUpdateResponseBody) *OnsGroupConsumerUpdateResponse
	GetBody() *OnsGroupConsumerUpdateResponseBody
}

type OnsGroupConsumerUpdateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupConsumerUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupConsumerUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupConsumerUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupConsumerUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsGroupConsumerUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsGroupConsumerUpdateResponse) GetBody() *OnsGroupConsumerUpdateResponseBody {
	return s.Body
}

func (s *OnsGroupConsumerUpdateResponse) SetHeaders(v map[string]*string) *OnsGroupConsumerUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupConsumerUpdateResponse) SetStatusCode(v int32) *OnsGroupConsumerUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupConsumerUpdateResponse) SetBody(v *OnsGroupConsumerUpdateResponseBody) *OnsGroupConsumerUpdateResponse {
	s.Body = v
	return s
}

func (s *OnsGroupConsumerUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
