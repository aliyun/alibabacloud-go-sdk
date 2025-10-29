// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLikeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLikeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLikeResponse
	GetStatusCode() *int32
	SetBody(v *SendLikeResponseBody) *SendLikeResponse
	GetBody() *SendLikeResponseBody
}

type SendLikeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLikeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLikeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLikeResponse) GoString() string {
	return s.String()
}

func (s *SendLikeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLikeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLikeResponse) GetBody() *SendLikeResponseBody {
	return s.Body
}

func (s *SendLikeResponse) SetHeaders(v map[string]*string) *SendLikeResponse {
	s.Headers = v
	return s
}

func (s *SendLikeResponse) SetStatusCode(v int32) *SendLikeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLikeResponse) SetBody(v *SendLikeResponseBody) *SendLikeResponse {
	s.Body = v
	return s
}

func (s *SendLikeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
