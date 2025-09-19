// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupIdByGroupNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGroupIdByGroupNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGroupIdByGroupNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryGroupIdByGroupNameResponseBody) *QueryGroupIdByGroupNameResponse
	GetBody() *QueryGroupIdByGroupNameResponseBody
}

type QueryGroupIdByGroupNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupIdByGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupIdByGroupNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupIdByGroupNameResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupIdByGroupNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGroupIdByGroupNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGroupIdByGroupNameResponse) GetBody() *QueryGroupIdByGroupNameResponseBody {
	return s.Body
}

func (s *QueryGroupIdByGroupNameResponse) SetHeaders(v map[string]*string) *QueryGroupIdByGroupNameResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupIdByGroupNameResponse) SetStatusCode(v int32) *QueryGroupIdByGroupNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupIdByGroupNameResponse) SetBody(v *QueryGroupIdByGroupNameResponseBody) *QueryGroupIdByGroupNameResponse {
	s.Body = v
	return s
}

func (s *QueryGroupIdByGroupNameResponse) Validate() error {
	return dara.Validate(s)
}
