// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcdpGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcdpGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcdpGroupResponseBody) *CreateMcdpGroupResponse
	GetBody() *CreateMcdpGroupResponseBody
}

type CreateMcdpGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcdpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcdpGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMcdpGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcdpGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcdpGroupResponse) GetBody() *CreateMcdpGroupResponseBody {
	return s.Body
}

func (s *CreateMcdpGroupResponse) SetHeaders(v map[string]*string) *CreateMcdpGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMcdpGroupResponse) SetStatusCode(v int32) *CreateMcdpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcdpGroupResponse) SetBody(v *CreateMcdpGroupResponseBody) *CreateMcdpGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMcdpGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
