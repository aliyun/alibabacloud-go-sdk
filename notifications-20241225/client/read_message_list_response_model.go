// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMessageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMessageListResponse
	GetStatusCode() *int32
	SetBody(v *ReadMessageListResponseBody) *ReadMessageListResponse
	GetBody() *ReadMessageListResponseBody
}

type ReadMessageListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageListResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMessageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMessageListResponse) GetBody() *ReadMessageListResponseBody {
	return s.Body
}

func (s *ReadMessageListResponse) SetHeaders(v map[string]*string) *ReadMessageListResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageListResponse) SetStatusCode(v int32) *ReadMessageListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageListResponse) SetBody(v *ReadMessageListResponseBody) *ReadMessageListResponse {
	s.Body = v
	return s
}

func (s *ReadMessageListResponse) Validate() error {
	return dara.Validate(s)
}
