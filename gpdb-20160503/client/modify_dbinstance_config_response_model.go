// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceConfigResponseBody) *ModifyDBInstanceConfigResponse
	GetBody() *ModifyDBInstanceConfigResponseBody
}

type ModifyDBInstanceConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceConfigResponse) GetBody() *ModifyDBInstanceConfigResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetStatusCode(v int32) *ModifyDBInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetBody(v *ModifyDBInstanceConfigResponseBody) *ModifyDBInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
