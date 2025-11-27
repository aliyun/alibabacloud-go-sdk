// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicListResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicListResponseBody) *GetTopicListResponse
	GetBody() *GetTopicListResponseBody
}

type GetTopicListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponse) GoString() string {
	return s.String()
}

func (s *GetTopicListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicListResponse) GetBody() *GetTopicListResponseBody {
	return s.Body
}

func (s *GetTopicListResponse) SetHeaders(v map[string]*string) *GetTopicListResponse {
	s.Headers = v
	return s
}

func (s *GetTopicListResponse) SetStatusCode(v int32) *GetTopicListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicListResponse) SetBody(v *GetTopicListResponseBody) *GetTopicListResponse {
	s.Body = v
	return s
}

func (s *GetTopicListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
