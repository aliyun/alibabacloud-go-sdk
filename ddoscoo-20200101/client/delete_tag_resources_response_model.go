// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTagResourcesResponseBody) *DeleteTagResourcesResponse
	GetBody() *DeleteTagResourcesResponseBody
}

type DeleteTagResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTagResourcesResponse) GetBody() *DeleteTagResourcesResponseBody {
	return s.Body
}

func (s *DeleteTagResourcesResponse) SetHeaders(v map[string]*string) *DeleteTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResourcesResponse) SetStatusCode(v int32) *DeleteTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagResourcesResponse) SetBody(v *DeleteTagResourcesResponseBody) *DeleteTagResourcesResponse {
	s.Body = v
	return s
}

func (s *DeleteTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
