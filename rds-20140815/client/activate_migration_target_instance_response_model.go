// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMigrationTargetInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateMigrationTargetInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateMigrationTargetInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ActivateMigrationTargetInstanceResponseBody) *ActivateMigrationTargetInstanceResponse
	GetBody() *ActivateMigrationTargetInstanceResponseBody
}

type ActivateMigrationTargetInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateMigrationTargetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateMigrationTargetInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateMigrationTargetInstanceResponse) GoString() string {
	return s.String()
}

func (s *ActivateMigrationTargetInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateMigrationTargetInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateMigrationTargetInstanceResponse) GetBody() *ActivateMigrationTargetInstanceResponseBody {
	return s.Body
}

func (s *ActivateMigrationTargetInstanceResponse) SetHeaders(v map[string]*string) *ActivateMigrationTargetInstanceResponse {
	s.Headers = v
	return s
}

func (s *ActivateMigrationTargetInstanceResponse) SetStatusCode(v int32) *ActivateMigrationTargetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateMigrationTargetInstanceResponse) SetBody(v *ActivateMigrationTargetInstanceResponseBody) *ActivateMigrationTargetInstanceResponse {
	s.Body = v
	return s
}

func (s *ActivateMigrationTargetInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
