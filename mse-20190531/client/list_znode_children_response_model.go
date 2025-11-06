// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZnodeChildrenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListZnodeChildrenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListZnodeChildrenResponse
	GetStatusCode() *int32
	SetBody(v *ListZnodeChildrenResponseBody) *ListZnodeChildrenResponse
	GetBody() *ListZnodeChildrenResponseBody
}

type ListZnodeChildrenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZnodeChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZnodeChildrenResponse) String() string {
	return dara.Prettify(s)
}

func (s ListZnodeChildrenResponse) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListZnodeChildrenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListZnodeChildrenResponse) GetBody() *ListZnodeChildrenResponseBody {
	return s.Body
}

func (s *ListZnodeChildrenResponse) SetHeaders(v map[string]*string) *ListZnodeChildrenResponse {
	s.Headers = v
	return s
}

func (s *ListZnodeChildrenResponse) SetStatusCode(v int32) *ListZnodeChildrenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZnodeChildrenResponse) SetBody(v *ListZnodeChildrenResponseBody) *ListZnodeChildrenResponse {
	s.Body = v
	return s
}

func (s *ListZnodeChildrenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
