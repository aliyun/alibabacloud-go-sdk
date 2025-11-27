// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceForRebuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceForRebuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceForRebuildResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceForRebuildResponseBody) *CreateDBInstanceForRebuildResponse
	GetBody() *CreateDBInstanceForRebuildResponseBody
}

type CreateDBInstanceForRebuildResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceForRebuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceForRebuildResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceForRebuildResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceForRebuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceForRebuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceForRebuildResponse) GetBody() *CreateDBInstanceForRebuildResponseBody {
	return s.Body
}

func (s *CreateDBInstanceForRebuildResponse) SetHeaders(v map[string]*string) *CreateDBInstanceForRebuildResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceForRebuildResponse) SetStatusCode(v int32) *CreateDBInstanceForRebuildResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceForRebuildResponse) SetBody(v *CreateDBInstanceForRebuildResponseBody) *CreateDBInstanceForRebuildResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceForRebuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
