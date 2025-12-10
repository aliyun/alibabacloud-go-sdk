// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateExperimentFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateExperimentFoldersResponse
	GetStatusCode() *int32
	SetBody(v *MigrateExperimentFoldersResponseBody) *MigrateExperimentFoldersResponse
	GetBody() *MigrateExperimentFoldersResponseBody
}

type MigrateExperimentFoldersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateExperimentFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateExperimentFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentFoldersResponse) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateExperimentFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateExperimentFoldersResponse) GetBody() *MigrateExperimentFoldersResponseBody {
	return s.Body
}

func (s *MigrateExperimentFoldersResponse) SetHeaders(v map[string]*string) *MigrateExperimentFoldersResponse {
	s.Headers = v
	return s
}

func (s *MigrateExperimentFoldersResponse) SetStatusCode(v int32) *MigrateExperimentFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateExperimentFoldersResponse) SetBody(v *MigrateExperimentFoldersResponseBody) *MigrateExperimentFoldersResponse {
	s.Body = v
	return s
}

func (s *MigrateExperimentFoldersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
