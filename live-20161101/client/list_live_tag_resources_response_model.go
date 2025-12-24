// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveTagResourcesResponseBody) *ListLiveTagResourcesResponse
	GetBody() *ListLiveTagResourcesResponseBody
}

type ListLiveTagResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveTagResourcesResponse) GetBody() *ListLiveTagResourcesResponseBody {
	return s.Body
}

func (s *ListLiveTagResourcesResponse) SetHeaders(v map[string]*string) *ListLiveTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveTagResourcesResponse) SetStatusCode(v int32) *ListLiveTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveTagResourcesResponse) SetBody(v *ListLiveTagResourcesResponseBody) *ListLiveTagResourcesResponse {
	s.Body = v
	return s
}

func (s *ListLiveTagResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
