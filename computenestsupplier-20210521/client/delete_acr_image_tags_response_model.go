// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAcrImageTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAcrImageTagsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAcrImageTagsResponseBody) *DeleteAcrImageTagsResponse
	GetBody() *DeleteAcrImageTagsResponseBody
}

type DeleteAcrImageTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAcrImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAcrImageTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAcrImageTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAcrImageTagsResponse) GetBody() *DeleteAcrImageTagsResponseBody {
	return s.Body
}

func (s *DeleteAcrImageTagsResponse) SetHeaders(v map[string]*string) *DeleteAcrImageTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAcrImageTagsResponse) SetStatusCode(v int32) *DeleteAcrImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAcrImageTagsResponse) SetBody(v *DeleteAcrImageTagsResponseBody) *DeleteAcrImageTagsResponse {
	s.Body = v
	return s
}

func (s *DeleteAcrImageTagsResponse) Validate() error {
	return dara.Validate(s)
}
