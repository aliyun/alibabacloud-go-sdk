// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBrowserInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBrowserInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListBrowserInstanceGroupResponseBody) *ListBrowserInstanceGroupResponse
	GetBody() *ListBrowserInstanceGroupResponseBody
}

type ListBrowserInstanceGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBrowserInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBrowserInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBrowserInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBrowserInstanceGroupResponse) GetBody() *ListBrowserInstanceGroupResponseBody {
	return s.Body
}

func (s *ListBrowserInstanceGroupResponse) SetHeaders(v map[string]*string) *ListBrowserInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListBrowserInstanceGroupResponse) SetStatusCode(v int32) *ListBrowserInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBrowserInstanceGroupResponse) SetBody(v *ListBrowserInstanceGroupResponseBody) *ListBrowserInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *ListBrowserInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
