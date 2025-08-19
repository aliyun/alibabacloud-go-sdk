// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagResourcesResponse
	GetStatusCode() *int32
}

type TagResourcesResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
