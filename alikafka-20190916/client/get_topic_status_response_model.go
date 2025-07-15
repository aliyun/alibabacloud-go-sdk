// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicStatusResponseBody) *GetTopicStatusResponse
	GetBody() *GetTopicStatusResponseBody
}

type GetTopicStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicStatusResponse) GetBody() *GetTopicStatusResponseBody {
	return s.Body
}

func (s *GetTopicStatusResponse) SetHeaders(v map[string]*string) *GetTopicStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTopicStatusResponse) SetStatusCode(v int32) *GetTopicStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicStatusResponse) SetBody(v *GetTopicStatusResponseBody) *GetTopicStatusResponse {
	s.Body = v
	return s
}

func (s *GetTopicStatusResponse) Validate() error {
	return dara.Validate(s)
}
