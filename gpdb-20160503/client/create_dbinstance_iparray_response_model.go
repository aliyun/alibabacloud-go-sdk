// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceIPArrayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceIPArrayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceIPArrayResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceIPArrayResponseBody) *CreateDBInstanceIPArrayResponse
	GetBody() *CreateDBInstanceIPArrayResponseBody
}

type CreateDBInstanceIPArrayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceIPArrayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceIPArrayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceIPArrayResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceIPArrayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceIPArrayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceIPArrayResponse) GetBody() *CreateDBInstanceIPArrayResponseBody {
	return s.Body
}

func (s *CreateDBInstanceIPArrayResponse) SetHeaders(v map[string]*string) *CreateDBInstanceIPArrayResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceIPArrayResponse) SetStatusCode(v int32) *CreateDBInstanceIPArrayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceIPArrayResponse) SetBody(v *CreateDBInstanceIPArrayResponseBody) *CreateDBInstanceIPArrayResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceIPArrayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
