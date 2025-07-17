// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardGroupResponseBody) *CreateStandardGroupResponse
	GetBody() *CreateStandardGroupResponseBody
}

type CreateStandardGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardGroupResponse) GetBody() *CreateStandardGroupResponseBody {
	return s.Body
}

func (s *CreateStandardGroupResponse) SetHeaders(v map[string]*string) *CreateStandardGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardGroupResponse) SetStatusCode(v int32) *CreateStandardGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardGroupResponse) SetBody(v *CreateStandardGroupResponseBody) *CreateStandardGroupResponse {
	s.Body = v
	return s
}

func (s *CreateStandardGroupResponse) Validate() error {
	return dara.Validate(s)
}
