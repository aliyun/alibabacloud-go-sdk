// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBrowserInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBrowserInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBrowserInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBrowserInstanceGroupResponseBody) *ModifyBrowserInstanceGroupResponse
	GetBody() *ModifyBrowserInstanceGroupResponseBody
}

type ModifyBrowserInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBrowserInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBrowserInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBrowserInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBrowserInstanceGroupResponse) GetBody() *ModifyBrowserInstanceGroupResponseBody {
	return s.Body
}

func (s *ModifyBrowserInstanceGroupResponse) SetHeaders(v map[string]*string) *ModifyBrowserInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyBrowserInstanceGroupResponse) SetStatusCode(v int32) *ModifyBrowserInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBrowserInstanceGroupResponse) SetBody(v *ModifyBrowserInstanceGroupResponseBody) *ModifyBrowserInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyBrowserInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
