// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicByIdResponseBody) *GetTopicByIdResponse
	GetBody() *GetTopicByIdResponseBody
}

type GetTopicByIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicByIdResponse) GetBody() *GetTopicByIdResponseBody {
	return s.Body
}

func (s *GetTopicByIdResponse) SetHeaders(v map[string]*string) *GetTopicByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTopicByIdResponse) SetStatusCode(v int32) *GetTopicByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicByIdResponse) SetBody(v *GetTopicByIdResponseBody) *GetTopicByIdResponse {
	s.Body = v
	return s
}

func (s *GetTopicByIdResponse) Validate() error {
	return dara.Validate(s)
}
