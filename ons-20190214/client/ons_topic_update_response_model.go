// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicUpdateResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicUpdateResponseBody) *OnsTopicUpdateResponse
	GetBody() *OnsTopicUpdateResponseBody
}

type OnsTopicUpdateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicUpdateResponse) GetBody() *OnsTopicUpdateResponseBody {
	return s.Body
}

func (s *OnsTopicUpdateResponse) SetHeaders(v map[string]*string) *OnsTopicUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicUpdateResponse) SetStatusCode(v int32) *OnsTopicUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicUpdateResponse) SetBody(v *OnsTopicUpdateResponseBody) *OnsTopicUpdateResponse {
	s.Body = v
	return s
}

func (s *OnsTopicUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
