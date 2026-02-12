// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicCreateResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicCreateResponseBody) *OnsTopicCreateResponse
	GetBody() *OnsTopicCreateResponseBody
}

type OnsTopicCreateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicCreateResponse) GetBody() *OnsTopicCreateResponseBody {
	return s.Body
}

func (s *OnsTopicCreateResponse) SetHeaders(v map[string]*string) *OnsTopicCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicCreateResponse) SetStatusCode(v int32) *OnsTopicCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicCreateResponse) SetBody(v *OnsTopicCreateResponseBody) *OnsTopicCreateResponse {
	s.Body = v
	return s
}

func (s *OnsTopicCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
