// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceResourceGroupResponseBody) *ModifyDBInstanceResourceGroupResponse
	GetBody() *ModifyDBInstanceResourceGroupResponseBody
}

type ModifyDBInstanceResourceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceResourceGroupResponse) GetBody() *ModifyDBInstanceResourceGroupResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetStatusCode(v int32) *ModifyDBInstanceResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetBody(v *ModifyDBInstanceResourceGroupResponseBody) *ModifyDBInstanceResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
