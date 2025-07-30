// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateParameterGroupResponseBody) *CreateParameterGroupResponse
	GetBody() *CreateParameterGroupResponseBody
}

type CreateParameterGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateParameterGroupResponse) GetBody() *CreateParameterGroupResponseBody {
	return s.Body
}

func (s *CreateParameterGroupResponse) SetHeaders(v map[string]*string) *CreateParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterGroupResponse) SetStatusCode(v int32) *CreateParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParameterGroupResponse) SetBody(v *CreateParameterGroupResponseBody) *CreateParameterGroupResponse {
	s.Body = v
	return s
}

func (s *CreateParameterGroupResponse) Validate() error {
	return dara.Validate(s)
}
