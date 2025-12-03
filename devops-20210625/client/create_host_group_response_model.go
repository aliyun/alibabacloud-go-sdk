// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHostGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHostGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse
	GetBody() *CreateHostGroupResponseBody
}

type CreateHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHostGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHostGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHostGroupResponse) GetBody() *CreateHostGroupResponseBody {
	return s.Body
}

func (s *CreateHostGroupResponse) SetHeaders(v map[string]*string) *CreateHostGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateHostGroupResponse) SetStatusCode(v int32) *CreateHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostGroupResponse) SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse {
	s.Body = v
	return s
}

func (s *CreateHostGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
