// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBResourceGroupResponseBody) *CreateDBResourceGroupResponse
	GetBody() *CreateDBResourceGroupResponseBody
}

type CreateDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBResourceGroupResponse) GetBody() *CreateDBResourceGroupResponseBody {
	return s.Body
}

func (s *CreateDBResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourceGroupResponse) SetStatusCode(v int32) *CreateDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResourceGroupResponse) SetBody(v *CreateDBResourceGroupResponseBody) *CreateDBResourceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDBResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
