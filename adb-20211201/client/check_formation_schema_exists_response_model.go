// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFormationSchemaExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckFormationSchemaExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckFormationSchemaExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckFormationSchemaExistsResponseBody) *CheckFormationSchemaExistsResponse
	GetBody() *CheckFormationSchemaExistsResponseBody
}

type CheckFormationSchemaExistsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFormationSchemaExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckFormationSchemaExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckFormationSchemaExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckFormationSchemaExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckFormationSchemaExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckFormationSchemaExistsResponse) GetBody() *CheckFormationSchemaExistsResponseBody {
	return s.Body
}

func (s *CheckFormationSchemaExistsResponse) SetHeaders(v map[string]*string) *CheckFormationSchemaExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckFormationSchemaExistsResponse) SetStatusCode(v int32) *CheckFormationSchemaExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFormationSchemaExistsResponse) SetBody(v *CheckFormationSchemaExistsResponseBody) *CheckFormationSchemaExistsResponse {
	s.Body = v
	return s
}

func (s *CheckFormationSchemaExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
