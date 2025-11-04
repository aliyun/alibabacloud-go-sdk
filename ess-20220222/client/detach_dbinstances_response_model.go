// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDBInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDBInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDBInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DetachDBInstancesResponseBody) *DetachDBInstancesResponse
	GetBody() *DetachDBInstancesResponseBody
}

type DetachDBInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDBInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDBInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDBInstancesResponse) GetBody() *DetachDBInstancesResponseBody {
	return s.Body
}

func (s *DetachDBInstancesResponse) SetHeaders(v map[string]*string) *DetachDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DetachDBInstancesResponse) SetStatusCode(v int32) *DetachDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDBInstancesResponse) SetBody(v *DetachDBInstancesResponseBody) *DetachDBInstancesResponse {
	s.Body = v
	return s
}

func (s *DetachDBInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
