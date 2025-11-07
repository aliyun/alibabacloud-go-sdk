// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationGroupResponseBody) *CreateApplicationGroupResponse
	GetBody() *CreateApplicationGroupResponseBody
}

type CreateApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationGroupResponse) GetBody() *CreateApplicationGroupResponseBody {
	return s.Body
}

func (s *CreateApplicationGroupResponse) SetHeaders(v map[string]*string) *CreateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationGroupResponse) SetStatusCode(v int32) *CreateApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationGroupResponse) SetBody(v *CreateApplicationGroupResponseBody) *CreateApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
