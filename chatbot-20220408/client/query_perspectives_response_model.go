// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPerspectivesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPerspectivesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPerspectivesResponse
	GetStatusCode() *int32
	SetBody(v *QueryPerspectivesResponseBody) *QueryPerspectivesResponse
	GetBody() *QueryPerspectivesResponseBody
}

type QueryPerspectivesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPerspectivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPerspectivesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPerspectivesResponse) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPerspectivesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPerspectivesResponse) GetBody() *QueryPerspectivesResponseBody {
	return s.Body
}

func (s *QueryPerspectivesResponse) SetHeaders(v map[string]*string) *QueryPerspectivesResponse {
	s.Headers = v
	return s
}

func (s *QueryPerspectivesResponse) SetStatusCode(v int32) *QueryPerspectivesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPerspectivesResponse) SetBody(v *QueryPerspectivesResponseBody) *QueryPerspectivesResponse {
	s.Body = v
	return s
}

func (s *QueryPerspectivesResponse) Validate() error {
	return dara.Validate(s)
}
