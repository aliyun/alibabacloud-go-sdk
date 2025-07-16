// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse
	GetBody() *DeleteDBInstanceResponseBody
}

type DeleteDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceResponse) GetBody() *DeleteDBInstanceResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceResponse) SetStatusCode(v int32) *DeleteDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceResponse) Validate() error {
	return dara.Validate(s)
}
