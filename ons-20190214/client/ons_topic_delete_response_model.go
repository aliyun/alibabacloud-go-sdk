// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicDeleteResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicDeleteResponseBody) *OnsTopicDeleteResponse
	GetBody() *OnsTopicDeleteResponseBody
}

type OnsTopicDeleteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicDeleteResponse) GetBody() *OnsTopicDeleteResponseBody {
	return s.Body
}

func (s *OnsTopicDeleteResponse) SetHeaders(v map[string]*string) *OnsTopicDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicDeleteResponse) SetStatusCode(v int32) *OnsTopicDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicDeleteResponse) SetBody(v *OnsTopicDeleteResponseBody) *OnsTopicDeleteResponse {
	s.Body = v
	return s
}

func (s *OnsTopicDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
