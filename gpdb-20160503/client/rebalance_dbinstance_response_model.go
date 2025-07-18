// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebalanceDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebalanceDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebalanceDBInstanceResponseBody) *RebalanceDBInstanceResponse
	GetBody() *RebalanceDBInstanceResponseBody
}

type RebalanceDBInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebalanceDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebalanceDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebalanceDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebalanceDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebalanceDBInstanceResponse) GetBody() *RebalanceDBInstanceResponseBody {
	return s.Body
}

func (s *RebalanceDBInstanceResponse) SetHeaders(v map[string]*string) *RebalanceDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebalanceDBInstanceResponse) SetStatusCode(v int32) *RebalanceDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceDBInstanceResponse) SetBody(v *RebalanceDBInstanceResponseBody) *RebalanceDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RebalanceDBInstanceResponse) Validate() error {
	return dara.Validate(s)
}
