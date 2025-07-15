// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePostPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePostPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreatePostPayOrderResponseBody) *CreatePostPayOrderResponse
	GetBody() *CreatePostPayOrderResponseBody
}

type CreatePostPayOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePostPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePostPayOrderResponse) GetBody() *CreatePostPayOrderResponseBody {
	return s.Body
}

func (s *CreatePostPayOrderResponse) SetHeaders(v map[string]*string) *CreatePostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePostPayOrderResponse) SetStatusCode(v int32) *CreatePostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostPayOrderResponse) SetBody(v *CreatePostPayOrderResponseBody) *CreatePostPayOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePostPayOrderResponse) Validate() error {
	return dara.Validate(s)
}
