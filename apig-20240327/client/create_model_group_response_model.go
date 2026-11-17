// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelGroupResponseBody) *CreateModelGroupResponse
	GetBody() *CreateModelGroupResponseBody
}

type CreateModelGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelGroupResponse) GetBody() *CreateModelGroupResponseBody {
	return s.Body
}

func (s *CreateModelGroupResponse) SetHeaders(v map[string]*string) *CreateModelGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateModelGroupResponse) SetStatusCode(v int32) *CreateModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelGroupResponse) SetBody(v *CreateModelGroupResponseBody) *CreateModelGroupResponse {
	s.Body = v
	return s
}

func (s *CreateModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
