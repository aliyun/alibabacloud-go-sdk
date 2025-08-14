// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivePackageChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivePackageChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListLivePackageChannelsResponseBody) *ListLivePackageChannelsResponse
	GetBody() *ListLivePackageChannelsResponseBody
}

type ListLivePackageChannelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivePackageChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivePackageChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivePackageChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivePackageChannelsResponse) GetBody() *ListLivePackageChannelsResponseBody {
	return s.Body
}

func (s *ListLivePackageChannelsResponse) SetHeaders(v map[string]*string) *ListLivePackageChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListLivePackageChannelsResponse) SetStatusCode(v int32) *ListLivePackageChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivePackageChannelsResponse) SetBody(v *ListLivePackageChannelsResponseBody) *ListLivePackageChannelsResponse {
	s.Body = v
	return s
}

func (s *ListLivePackageChannelsResponse) Validate() error {
	return dara.Validate(s)
}
