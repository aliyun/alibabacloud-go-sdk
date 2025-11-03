// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceTDEResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceTDEResponseBody) *ModifyDBInstanceTDEResponse
	GetBody() *ModifyDBInstanceTDEResponseBody
}

type ModifyDBInstanceTDEResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceTDEResponse) GetBody() *ModifyDBInstanceTDEResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceTDEResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceTDEResponse) SetStatusCode(v int32) *ModifyDBInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceTDEResponse) SetBody(v *ModifyDBInstanceTDEResponseBody) *ModifyDBInstanceTDEResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceTDEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
