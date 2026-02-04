// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceCLSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceCLSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceCLSResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceCLSResponseBody) *ModifyDBInstanceCLSResponse
	GetBody() *ModifyDBInstanceCLSResponseBody
}

type ModifyDBInstanceCLSResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceCLSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceCLSResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceCLSResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceCLSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceCLSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceCLSResponse) GetBody() *ModifyDBInstanceCLSResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceCLSResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceCLSResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceCLSResponse) SetStatusCode(v int32) *ModifyDBInstanceCLSResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceCLSResponse) SetBody(v *ModifyDBInstanceCLSResponseBody) *ModifyDBInstanceCLSResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceCLSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
