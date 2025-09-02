// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEntityTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetEntityTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetEntityTagsResponse
	GetStatusCode() *int32
	SetBody(v *SetEntityTagsResponseBody) *SetEntityTagsResponse
	GetBody() *SetEntityTagsResponseBody
}

type SetEntityTagsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEntityTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEntityTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetEntityTagsResponse) GoString() string {
	return s.String()
}

func (s *SetEntityTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetEntityTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetEntityTagsResponse) GetBody() *SetEntityTagsResponseBody {
	return s.Body
}

func (s *SetEntityTagsResponse) SetHeaders(v map[string]*string) *SetEntityTagsResponse {
	s.Headers = v
	return s
}

func (s *SetEntityTagsResponse) SetStatusCode(v int32) *SetEntityTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEntityTagsResponse) SetBody(v *SetEntityTagsResponseBody) *SetEntityTagsResponse {
	s.Body = v
	return s
}

func (s *SetEntityTagsResponse) Validate() error {
	return dara.Validate(s)
}
