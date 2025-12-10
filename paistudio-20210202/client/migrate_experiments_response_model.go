// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateExperimentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateExperimentsResponse
	GetStatusCode() *int32
	SetBody(v *MigrateExperimentsResponseBody) *MigrateExperimentsResponse
	GetBody() *MigrateExperimentsResponseBody
}

type MigrateExperimentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateExperimentsResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentsResponse) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateExperimentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateExperimentsResponse) GetBody() *MigrateExperimentsResponseBody {
	return s.Body
}

func (s *MigrateExperimentsResponse) SetHeaders(v map[string]*string) *MigrateExperimentsResponse {
	s.Headers = v
	return s
}

func (s *MigrateExperimentsResponse) SetStatusCode(v int32) *MigrateExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateExperimentsResponse) SetBody(v *MigrateExperimentsResponseBody) *MigrateExperimentsResponse {
	s.Body = v
	return s
}

func (s *MigrateExperimentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
