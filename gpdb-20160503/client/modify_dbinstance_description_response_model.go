// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse
	GetBody() *ModifyDBInstanceDescriptionResponseBody
}

type ModifyDBInstanceDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceDescriptionResponse) GetBody() *ModifyDBInstanceDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDBInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
