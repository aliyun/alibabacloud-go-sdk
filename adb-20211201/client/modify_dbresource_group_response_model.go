// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBResourceGroupResponseBody) *ModifyDBResourceGroupResponse
	GetBody() *ModifyDBResourceGroupResponseBody
}

type ModifyDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBResourceGroupResponse) GetBody() *ModifyDBResourceGroupResponseBody {
	return s.Body
}

func (s *ModifyDBResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetStatusCode(v int32) *ModifyDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetBody(v *ModifyDBResourceGroupResponseBody) *ModifyDBResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDBResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
