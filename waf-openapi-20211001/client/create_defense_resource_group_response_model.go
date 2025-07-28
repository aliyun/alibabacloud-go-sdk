// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefenseResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefenseResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefenseResourceGroupResponseBody) *CreateDefenseResourceGroupResponse
	GetBody() *CreateDefenseResourceGroupResponseBody
}

type CreateDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefenseResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefenseResourceGroupResponse) GetBody() *CreateDefenseResourceGroupResponseBody {
	return s.Body
}

func (s *CreateDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetStatusCode(v int32) *CreateDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetBody(v *CreateDefenseResourceGroupResponseBody) *CreateDefenseResourceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDefenseResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
