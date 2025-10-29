// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagLiveResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnTagLiveResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnTagLiveResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UnTagLiveResourcesResponseBody) *UnTagLiveResourcesResponse
	GetBody() *UnTagLiveResourcesResponseBody
}

type UnTagLiveResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagLiveResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagLiveResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnTagLiveResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagLiveResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnTagLiveResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnTagLiveResourcesResponse) GetBody() *UnTagLiveResourcesResponseBody {
	return s.Body
}

func (s *UnTagLiveResourcesResponse) SetHeaders(v map[string]*string) *UnTagLiveResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagLiveResourcesResponse) SetStatusCode(v int32) *UnTagLiveResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagLiveResourcesResponse) SetBody(v *UnTagLiveResourcesResponseBody) *UnTagLiveResourcesResponse {
	s.Body = v
	return s
}

func (s *UnTagLiveResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
