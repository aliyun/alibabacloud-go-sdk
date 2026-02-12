// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicListResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicListResponseBody) *OnsTopicListResponse
	GetBody() *OnsTopicListResponseBody
}

type OnsTopicListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicListResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicListResponse) GetBody() *OnsTopicListResponseBody {
	return s.Body
}

func (s *OnsTopicListResponse) SetHeaders(v map[string]*string) *OnsTopicListResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicListResponse) SetStatusCode(v int32) *OnsTopicListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicListResponse) SetBody(v *OnsTopicListResponseBody) *OnsTopicListResponse {
	s.Body = v
	return s
}

func (s *OnsTopicListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
