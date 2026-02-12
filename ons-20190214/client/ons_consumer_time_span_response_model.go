// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerTimeSpanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsConsumerTimeSpanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsConsumerTimeSpanResponse
	GetStatusCode() *int32
	SetBody(v *OnsConsumerTimeSpanResponseBody) *OnsConsumerTimeSpanResponse
	GetBody() *OnsConsumerTimeSpanResponseBody
}

type OnsConsumerTimeSpanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerTimeSpanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerTimeSpanResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerTimeSpanResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsConsumerTimeSpanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsConsumerTimeSpanResponse) GetBody() *OnsConsumerTimeSpanResponseBody {
	return s.Body
}

func (s *OnsConsumerTimeSpanResponse) SetHeaders(v map[string]*string) *OnsConsumerTimeSpanResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerTimeSpanResponse) SetStatusCode(v int32) *OnsConsumerTimeSpanResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerTimeSpanResponse) SetBody(v *OnsConsumerTimeSpanResponseBody) *OnsConsumerTimeSpanResponse {
	s.Body = v
	return s
}

func (s *OnsConsumerTimeSpanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
