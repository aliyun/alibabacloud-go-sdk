// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBResourceGroupResponseBody) *DisableDBResourceGroupResponse
	GetBody() *DisableDBResourceGroupResponseBody
}

type DisableDBResourceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableDBResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBResourceGroupResponse) GetBody() *DisableDBResourceGroupResponseBody {
	return s.Body
}

func (s *DisableDBResourceGroupResponse) SetHeaders(v map[string]*string) *DisableDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableDBResourceGroupResponse) SetStatusCode(v int32) *DisableDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBResourceGroupResponse) SetBody(v *DisableDBResourceGroupResponseBody) *DisableDBResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DisableDBResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
