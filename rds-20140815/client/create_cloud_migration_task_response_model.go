// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudMigrationTaskResponseBody) *CreateCloudMigrationTaskResponse
	GetBody() *CreateCloudMigrationTaskResponseBody
}

type CreateCloudMigrationTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudMigrationTaskResponse) GetBody() *CreateCloudMigrationTaskResponseBody {
	return s.Body
}

func (s *CreateCloudMigrationTaskResponse) SetHeaders(v map[string]*string) *CreateCloudMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudMigrationTaskResponse) SetStatusCode(v int32) *CreateCloudMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudMigrationTaskResponse) SetBody(v *CreateCloudMigrationTaskResponseBody) *CreateCloudMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCloudMigrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
