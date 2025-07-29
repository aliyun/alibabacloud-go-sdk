// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse
	GetBody() *CreateAppGroupResponseBody
}

type CreateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppGroupResponse) GetBody() *CreateAppGroupResponseBody {
	return s.Body
}

func (s *CreateAppGroupResponse) SetHeaders(v map[string]*string) *CreateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupResponse) SetStatusCode(v int32) *CreateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAppGroupResponse) Validate() error {
	return dara.Validate(s)
}
