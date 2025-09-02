// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveEntityTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveEntityTagsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveEntityTagsResponseBody) *RemoveEntityTagsResponse
	GetBody() *RemoveEntityTagsResponseBody
}

type RemoveEntityTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEntityTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEntityTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntityTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveEntityTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveEntityTagsResponse) GetBody() *RemoveEntityTagsResponseBody {
	return s.Body
}

func (s *RemoveEntityTagsResponse) SetHeaders(v map[string]*string) *RemoveEntityTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntityTagsResponse) SetStatusCode(v int32) *RemoveEntityTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntityTagsResponse) SetBody(v *RemoveEntityTagsResponseBody) *RemoveEntityTagsResponse {
	s.Body = v
	return s
}

func (s *RemoveEntityTagsResponse) Validate() error {
	return dara.Validate(s)
}
