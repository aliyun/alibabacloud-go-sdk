// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationSubmitInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBwmMigrationSubmitInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBwmMigrationSubmitInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetBwmMigrationSubmitInstanceListResponseBody) *GetBwmMigrationSubmitInstanceListResponse
	GetBody() *GetBwmMigrationSubmitInstanceListResponseBody
}

type GetBwmMigrationSubmitInstanceListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBwmMigrationSubmitInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBwmMigrationSubmitInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationSubmitInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationSubmitInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBwmMigrationSubmitInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBwmMigrationSubmitInstanceListResponse) GetBody() *GetBwmMigrationSubmitInstanceListResponseBody {
	return s.Body
}

func (s *GetBwmMigrationSubmitInstanceListResponse) SetHeaders(v map[string]*string) *GetBwmMigrationSubmitInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponse) SetStatusCode(v int32) *GetBwmMigrationSubmitInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponse) SetBody(v *GetBwmMigrationSubmitInstanceListResponseBody) *GetBwmMigrationSubmitInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
