// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationPrecheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudMigrationPrecheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudMigrationPrecheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudMigrationPrecheckTaskResponseBody) *CreateCloudMigrationPrecheckTaskResponse
	GetBody() *CreateCloudMigrationPrecheckTaskResponseBody
}

type CreateCloudMigrationPrecheckTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudMigrationPrecheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudMigrationPrecheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationPrecheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationPrecheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudMigrationPrecheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudMigrationPrecheckTaskResponse) GetBody() *CreateCloudMigrationPrecheckTaskResponseBody {
	return s.Body
}

func (s *CreateCloudMigrationPrecheckTaskResponse) SetHeaders(v map[string]*string) *CreateCloudMigrationPrecheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponse) SetStatusCode(v int32) *CreateCloudMigrationPrecheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponse) SetBody(v *CreateCloudMigrationPrecheckTaskResponseBody) *CreateCloudMigrationPrecheckTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
