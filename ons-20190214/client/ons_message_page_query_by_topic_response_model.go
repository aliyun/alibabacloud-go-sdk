// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePageQueryByTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessagePageQueryByTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessagePageQueryByTopicResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessagePageQueryByTopicResponseBody) *OnsMessagePageQueryByTopicResponse
	GetBody() *OnsMessagePageQueryByTopicResponseBody
}

type OnsMessagePageQueryByTopicResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessagePageQueryByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponse) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessagePageQueryByTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessagePageQueryByTopicResponse) GetBody() *OnsMessagePageQueryByTopicResponseBody {
	return s.Body
}

func (s *OnsMessagePageQueryByTopicResponse) SetHeaders(v map[string]*string) *OnsMessagePageQueryByTopicResponse {
	s.Headers = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponse) SetStatusCode(v int32) *OnsMessagePageQueryByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponse) SetBody(v *OnsMessagePageQueryByTopicResponseBody) *OnsMessagePageQueryByTopicResponse {
	s.Body = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
