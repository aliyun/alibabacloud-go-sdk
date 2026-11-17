// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeConsumerModelCardGrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeConsumerModelCardGrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeConsumerModelCardGrantsResponse
	GetStatusCode() *int32
	SetBody(v *RevokeConsumerModelCardGrantsResponseBody) *RevokeConsumerModelCardGrantsResponse
	GetBody() *RevokeConsumerModelCardGrantsResponseBody
}

type RevokeConsumerModelCardGrantsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeConsumerModelCardGrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeConsumerModelCardGrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeConsumerModelCardGrantsResponse) GoString() string {
	return s.String()
}

func (s *RevokeConsumerModelCardGrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeConsumerModelCardGrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeConsumerModelCardGrantsResponse) GetBody() *RevokeConsumerModelCardGrantsResponseBody {
	return s.Body
}

func (s *RevokeConsumerModelCardGrantsResponse) SetHeaders(v map[string]*string) *RevokeConsumerModelCardGrantsResponse {
	s.Headers = v
	return s
}

func (s *RevokeConsumerModelCardGrantsResponse) SetStatusCode(v int32) *RevokeConsumerModelCardGrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeConsumerModelCardGrantsResponse) SetBody(v *RevokeConsumerModelCardGrantsResponseBody) *RevokeConsumerModelCardGrantsResponse {
	s.Body = v
	return s
}

func (s *RevokeConsumerModelCardGrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
