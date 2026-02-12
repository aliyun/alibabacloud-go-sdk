// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicSubDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTopicSubDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTopicSubDetailResponse
	GetStatusCode() *int32
	SetBody(v *OnsTopicSubDetailResponseBody) *OnsTopicSubDetailResponse
	GetBody() *OnsTopicSubDetailResponseBody
}

type OnsTopicSubDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicSubDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTopicSubDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTopicSubDetailResponse) GetBody() *OnsTopicSubDetailResponseBody {
	return s.Body
}

func (s *OnsTopicSubDetailResponse) SetHeaders(v map[string]*string) *OnsTopicSubDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicSubDetailResponse) SetStatusCode(v int32) *OnsTopicSubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicSubDetailResponse) SetBody(v *OnsTopicSubDetailResponseBody) *OnsTopicSubDetailResponse {
	s.Body = v
	return s
}

func (s *OnsTopicSubDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
