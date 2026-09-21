// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBrowserInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBrowserInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetBrowserInstanceGroupResponseBody) *GetBrowserInstanceGroupResponse
	GetBody() *GetBrowserInstanceGroupResponseBody
}

type GetBrowserInstanceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBrowserInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBrowserInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBrowserInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBrowserInstanceGroupResponse) GetBody() *GetBrowserInstanceGroupResponseBody {
	return s.Body
}

func (s *GetBrowserInstanceGroupResponse) SetHeaders(v map[string]*string) *GetBrowserInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetBrowserInstanceGroupResponse) SetStatusCode(v int32) *GetBrowserInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBrowserInstanceGroupResponse) SetBody(v *GetBrowserInstanceGroupResponseBody) *GetBrowserInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *GetBrowserInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
