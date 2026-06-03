// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDayuMigrateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDayuMigrateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDayuMigrateStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDayuMigrateStatusResponseBody) *QueryDayuMigrateStatusResponse
	GetBody() *QueryDayuMigrateStatusResponseBody
}

type QueryDayuMigrateStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDayuMigrateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDayuMigrateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDayuMigrateStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDayuMigrateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDayuMigrateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDayuMigrateStatusResponse) GetBody() *QueryDayuMigrateStatusResponseBody {
	return s.Body
}

func (s *QueryDayuMigrateStatusResponse) SetHeaders(v map[string]*string) *QueryDayuMigrateStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDayuMigrateStatusResponse) SetStatusCode(v int32) *QueryDayuMigrateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDayuMigrateStatusResponse) SetBody(v *QueryDayuMigrateStatusResponseBody) *QueryDayuMigrateStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDayuMigrateStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
