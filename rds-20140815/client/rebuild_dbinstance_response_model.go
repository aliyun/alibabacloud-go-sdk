// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebuildDBInstanceResponseBody) *RebuildDBInstanceResponse
	GetBody() *RebuildDBInstanceResponseBody
}

type RebuildDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebuildDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildDBInstanceResponse) GetBody() *RebuildDBInstanceResponseBody {
	return s.Body
}

func (s *RebuildDBInstanceResponse) SetHeaders(v map[string]*string) *RebuildDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebuildDBInstanceResponse) SetStatusCode(v int32) *RebuildDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildDBInstanceResponse) SetBody(v *RebuildDBInstanceResponseBody) *RebuildDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RebuildDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
