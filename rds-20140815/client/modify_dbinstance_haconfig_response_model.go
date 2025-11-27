// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceHAConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceHAConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceHAConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceHAConfigResponseBody) *ModifyDBInstanceHAConfigResponse
	GetBody() *ModifyDBInstanceHAConfigResponseBody
}

type ModifyDBInstanceHAConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceHAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceHAConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceHAConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceHAConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceHAConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceHAConfigResponse) GetBody() *ModifyDBInstanceHAConfigResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceHAConfigResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceHAConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceHAConfigResponse) SetStatusCode(v int32) *ModifyDBInstanceHAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceHAConfigResponse) SetBody(v *ModifyDBInstanceHAConfigResponseBody) *ModifyDBInstanceHAConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceHAConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
