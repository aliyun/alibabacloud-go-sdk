// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivePackageChannelGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivePackageChannelGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListLivePackageChannelGroupsResponseBody) *ListLivePackageChannelGroupsResponse
	GetBody() *ListLivePackageChannelGroupsResponseBody
}

type ListLivePackageChannelGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivePackageChannelGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivePackageChannelGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivePackageChannelGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivePackageChannelGroupsResponse) GetBody() *ListLivePackageChannelGroupsResponseBody {
	return s.Body
}

func (s *ListLivePackageChannelGroupsResponse) SetHeaders(v map[string]*string) *ListLivePackageChannelGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListLivePackageChannelGroupsResponse) SetStatusCode(v int32) *ListLivePackageChannelGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponse) SetBody(v *ListLivePackageChannelGroupsResponseBody) *ListLivePackageChannelGroupsResponse {
	s.Body = v
	return s
}

func (s *ListLivePackageChannelGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
