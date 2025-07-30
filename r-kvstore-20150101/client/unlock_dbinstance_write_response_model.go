// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDBInstanceWriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockDBInstanceWriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockDBInstanceWriteResponse
	GetStatusCode() *int32
	SetBody(v *UnlockDBInstanceWriteResponseBody) *UnlockDBInstanceWriteResponse
	GetBody() *UnlockDBInstanceWriteResponseBody
}

type UnlockDBInstanceWriteResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockDBInstanceWriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockDBInstanceWriteResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockDBInstanceWriteResponse) GoString() string {
	return s.String()
}

func (s *UnlockDBInstanceWriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockDBInstanceWriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockDBInstanceWriteResponse) GetBody() *UnlockDBInstanceWriteResponseBody {
	return s.Body
}

func (s *UnlockDBInstanceWriteResponse) SetHeaders(v map[string]*string) *UnlockDBInstanceWriteResponse {
	s.Headers = v
	return s
}

func (s *UnlockDBInstanceWriteResponse) SetStatusCode(v int32) *UnlockDBInstanceWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockDBInstanceWriteResponse) SetBody(v *UnlockDBInstanceWriteResponseBody) *UnlockDBInstanceWriteResponse {
	s.Body = v
	return s
}

func (s *UnlockDBInstanceWriteResponse) Validate() error {
	return dara.Validate(s)
}
