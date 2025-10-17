// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrationJobAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrationJobAlertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrationJobAlertResponseBody) *DescribeMigrationJobAlertResponse
	GetBody() *DescribeMigrationJobAlertResponseBody
}

type DescribeMigrationJobAlertResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrationJobAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrationJobAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrationJobAlertResponse) GetBody() *DescribeMigrationJobAlertResponseBody {
	return s.Body
}

func (s *DescribeMigrationJobAlertResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobAlertResponse) SetStatusCode(v int32) *DescribeMigrationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobAlertResponse) SetBody(v *DescribeMigrationJobAlertResponseBody) *DescribeMigrationJobAlertResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrationJobAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
