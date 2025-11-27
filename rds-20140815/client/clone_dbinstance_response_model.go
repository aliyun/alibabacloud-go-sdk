// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CloneDBInstanceResponseBody) *CloneDBInstanceResponse
	GetBody() *CloneDBInstanceResponseBody
}

type CloneDBInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneDBInstanceResponse) GetBody() *CloneDBInstanceResponseBody {
	return s.Body
}

func (s *CloneDBInstanceResponse) SetHeaders(v map[string]*string) *CloneDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CloneDBInstanceResponse) SetStatusCode(v int32) *CloneDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneDBInstanceResponse) SetBody(v *CloneDBInstanceResponseBody) *CloneDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CloneDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
