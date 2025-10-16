// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDesktopGroupResponseBody) *CreateDesktopGroupResponse
	GetBody() *CreateDesktopGroupResponseBody
}

type CreateDesktopGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDesktopGroupResponse) GetBody() *CreateDesktopGroupResponseBody {
	return s.Body
}

func (s *CreateDesktopGroupResponse) SetHeaders(v map[string]*string) *CreateDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopGroupResponse) SetStatusCode(v int32) *CreateDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDesktopGroupResponse) SetBody(v *CreateDesktopGroupResponseBody) *CreateDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDesktopGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
