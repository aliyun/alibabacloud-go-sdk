// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInfosForPodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeInfosForPodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeInfosForPodResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeInfosForPodResponseBody) *ListNodeInfosForPodResponse
	GetBody() *ListNodeInfosForPodResponseBody
}

type ListNodeInfosForPodResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeInfosForPodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeInfosForPodResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInfosForPodResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeInfosForPodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeInfosForPodResponse) GetBody() *ListNodeInfosForPodResponseBody {
	return s.Body
}

func (s *ListNodeInfosForPodResponse) SetHeaders(v map[string]*string) *ListNodeInfosForPodResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInfosForPodResponse) SetStatusCode(v int32) *ListNodeInfosForPodResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInfosForPodResponse) SetBody(v *ListNodeInfosForPodResponseBody) *ListNodeInfosForPodResponse {
	s.Body = v
	return s
}

func (s *ListNodeInfosForPodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
