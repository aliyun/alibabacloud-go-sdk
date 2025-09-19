// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreCheckItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IgnoreCheckItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IgnoreCheckItemsResponse
	GetStatusCode() *int32
	SetBody(v *IgnoreCheckItemsResponseBody) *IgnoreCheckItemsResponse
	GetBody() *IgnoreCheckItemsResponseBody
}

type IgnoreCheckItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IgnoreCheckItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IgnoreCheckItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s IgnoreCheckItemsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreCheckItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IgnoreCheckItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IgnoreCheckItemsResponse) GetBody() *IgnoreCheckItemsResponseBody {
	return s.Body
}

func (s *IgnoreCheckItemsResponse) SetHeaders(v map[string]*string) *IgnoreCheckItemsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreCheckItemsResponse) SetStatusCode(v int32) *IgnoreCheckItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreCheckItemsResponse) SetBody(v *IgnoreCheckItemsResponseBody) *IgnoreCheckItemsResponse {
	s.Body = v
	return s
}

func (s *IgnoreCheckItemsResponse) Validate() error {
	return dara.Validate(s)
}
