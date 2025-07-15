// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBgpGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBgpGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateBgpGroupResponseBody) *CreateBgpGroupResponse
	GetBody() *CreateBgpGroupResponseBody
}

type CreateBgpGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBgpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBgpGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateBgpGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBgpGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBgpGroupResponse) GetBody() *CreateBgpGroupResponseBody {
	return s.Body
}

func (s *CreateBgpGroupResponse) SetHeaders(v map[string]*string) *CreateBgpGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateBgpGroupResponse) SetStatusCode(v int32) *CreateBgpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBgpGroupResponse) SetBody(v *CreateBgpGroupResponseBody) *CreateBgpGroupResponse {
	s.Body = v
	return s
}

func (s *CreateBgpGroupResponse) Validate() error {
	return dara.Validate(s)
}
