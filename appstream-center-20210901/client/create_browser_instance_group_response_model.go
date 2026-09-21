// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBrowserInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBrowserInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateBrowserInstanceGroupResponseBody) *CreateBrowserInstanceGroupResponse
	GetBody() *CreateBrowserInstanceGroupResponseBody
}

type CreateBrowserInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBrowserInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBrowserInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBrowserInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBrowserInstanceGroupResponse) GetBody() *CreateBrowserInstanceGroupResponseBody {
	return s.Body
}

func (s *CreateBrowserInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateBrowserInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateBrowserInstanceGroupResponse) SetStatusCode(v int32) *CreateBrowserInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBrowserInstanceGroupResponse) SetBody(v *CreateBrowserInstanceGroupResponseBody) *CreateBrowserInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateBrowserInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
