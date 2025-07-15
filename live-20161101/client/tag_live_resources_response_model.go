// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagLiveResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagLiveResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagLiveResourcesResponse
	GetStatusCode() *int32
	SetBody(v *TagLiveResourcesResponseBody) *TagLiveResourcesResponse
	GetBody() *TagLiveResourcesResponseBody
}

type TagLiveResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagLiveResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagLiveResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagLiveResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagLiveResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagLiveResourcesResponse) GetBody() *TagLiveResourcesResponseBody {
	return s.Body
}

func (s *TagLiveResourcesResponse) SetHeaders(v map[string]*string) *TagLiveResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagLiveResourcesResponse) SetStatusCode(v int32) *TagLiveResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagLiveResourcesResponse) SetBody(v *TagLiveResourcesResponseBody) *TagLiveResourcesResponse {
	s.Body = v
	return s
}

func (s *TagLiveResourcesResponse) Validate() error {
	return dara.Validate(s)
}
