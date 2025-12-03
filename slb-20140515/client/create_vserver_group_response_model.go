// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateVServerGroupResponseBody) *CreateVServerGroupResponse
	GetBody() *CreateVServerGroupResponseBody
}

type CreateVServerGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVServerGroupResponse) GetBody() *CreateVServerGroupResponseBody {
	return s.Body
}

func (s *CreateVServerGroupResponse) SetHeaders(v map[string]*string) *CreateVServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVServerGroupResponse) SetStatusCode(v int32) *CreateVServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVServerGroupResponse) SetBody(v *CreateVServerGroupResponseBody) *CreateVServerGroupResponse {
	s.Body = v
	return s
}

func (s *CreateVServerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
