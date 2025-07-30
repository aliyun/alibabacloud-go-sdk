// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceAttributeResponseBody) *ModifyDBInstanceAttributeResponse
	GetBody() *ModifyDBInstanceAttributeResponseBody
}

type ModifyDBInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceAttributeResponse) GetBody() *ModifyDBInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceAttributeResponse) SetStatusCode(v int32) *ModifyDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponse) SetBody(v *ModifyDBInstanceAttributeResponseBody) *ModifyDBInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
