// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTaskByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrateTaskByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrateTaskByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrateTaskByIdResponseBody) *DescribeMigrateTaskByIdResponse
	GetBody() *DescribeMigrateTaskByIdResponseBody
}

type DescribeMigrateTaskByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrateTaskByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrateTaskByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTaskByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTaskByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrateTaskByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrateTaskByIdResponse) GetBody() *DescribeMigrateTaskByIdResponseBody {
	return s.Body
}

func (s *DescribeMigrateTaskByIdResponse) SetHeaders(v map[string]*string) *DescribeMigrateTaskByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrateTaskByIdResponse) SetStatusCode(v int32) *DescribeMigrateTaskByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponse) SetBody(v *DescribeMigrateTaskByIdResponseBody) *DescribeMigrateTaskByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrateTaskByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
