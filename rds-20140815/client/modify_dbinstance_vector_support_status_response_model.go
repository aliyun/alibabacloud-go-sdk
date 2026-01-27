// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVectorSupportStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceVectorSupportStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceVectorSupportStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceVectorSupportStatusResponseBody) *ModifyDBInstanceVectorSupportStatusResponse
	GetBody() *ModifyDBInstanceVectorSupportStatusResponseBody
}

type ModifyDBInstanceVectorSupportStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceVectorSupportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceVectorSupportStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVectorSupportStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) GetBody() *ModifyDBInstanceVectorSupportStatusResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceVectorSupportStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) SetStatusCode(v int32) *ModifyDBInstanceVectorSupportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) SetBody(v *ModifyDBInstanceVectorSupportStatusResponseBody) *ModifyDBInstanceVectorSupportStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
