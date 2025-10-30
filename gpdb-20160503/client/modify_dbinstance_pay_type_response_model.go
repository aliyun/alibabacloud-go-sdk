// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstancePayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstancePayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstancePayTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstancePayTypeResponseBody) *ModifyDBInstancePayTypeResponse
	GetBody() *ModifyDBInstancePayTypeResponseBody
}

type ModifyDBInstancePayTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstancePayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstancePayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstancePayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstancePayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstancePayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstancePayTypeResponse) GetBody() *ModifyDBInstancePayTypeResponseBody {
	return s.Body
}

func (s *ModifyDBInstancePayTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstancePayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstancePayTypeResponse) SetStatusCode(v int32) *ModifyDBInstancePayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstancePayTypeResponse) SetBody(v *ModifyDBInstancePayTypeResponseBody) *ModifyDBInstancePayTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstancePayTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
