// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReadOnlyDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReadOnlyDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReadOnlyDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateReadOnlyDBInstanceResponseBody) *CreateReadOnlyDBInstanceResponse
	GetBody() *CreateReadOnlyDBInstanceResponseBody
}

type CreateReadOnlyDBInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReadOnlyDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReadOnlyDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReadOnlyDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReadOnlyDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReadOnlyDBInstanceResponse) GetBody() *CreateReadOnlyDBInstanceResponseBody {
	return s.Body
}

func (s *CreateReadOnlyDBInstanceResponse) SetHeaders(v map[string]*string) *CreateReadOnlyDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateReadOnlyDBInstanceResponse) SetStatusCode(v int32) *CreateReadOnlyDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponse) SetBody(v *CreateReadOnlyDBInstanceResponseBody) *CreateReadOnlyDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateReadOnlyDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
