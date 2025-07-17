// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDBInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDBInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDBInstancesResponse
	GetStatusCode() *int32
	SetBody(v *AttachDBInstancesResponseBody) *AttachDBInstancesResponse
	GetBody() *AttachDBInstancesResponseBody
}

type AttachDBInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDBInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDBInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDBInstancesResponse) GetBody() *AttachDBInstancesResponseBody {
	return s.Body
}

func (s *AttachDBInstancesResponse) SetHeaders(v map[string]*string) *AttachDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachDBInstancesResponse) SetStatusCode(v int32) *AttachDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDBInstancesResponse) SetBody(v *AttachDBInstancesResponseBody) *AttachDBInstancesResponse {
	s.Body = v
	return s
}

func (s *AttachDBInstancesResponse) Validate() error {
	return dara.Validate(s)
}
