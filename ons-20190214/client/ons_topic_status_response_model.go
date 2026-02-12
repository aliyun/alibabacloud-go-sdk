// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicStatusResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicStatusResponseBody) *OnsTopicStatusResponse
	GetBody() *OnsTopicStatusResponseBody
}

type OnsTopicStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicStatusResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicStatusResponse) GetBody() *OnsTopicStatusResponseBody {
	return s.Body
}

func (s *OnsTopicStatusResponse) SetHeaders(v map[string]*string) *OnsTopicStatusResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicStatusResponse) SetStatusCode(v int32) *OnsTopicStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicStatusResponse) SetBody(v *OnsTopicStatusResponseBody) *OnsTopicStatusResponse {
	s.Body = v
	return s
}

func (s *OnsTopicStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
