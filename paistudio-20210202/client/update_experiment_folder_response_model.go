// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentFolderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentFolderResponseBody) *UpdateExperimentFolderResponse
	GetBody() *UpdateExperimentFolderResponseBody
}

type UpdateExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentFolderResponse) GetBody() *UpdateExperimentFolderResponseBody {
	return s.Body
}

func (s *UpdateExperimentFolderResponse) SetHeaders(v map[string]*string) *UpdateExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentFolderResponse) SetStatusCode(v int32) *UpdateExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentFolderResponse) SetBody(v *UpdateExperimentFolderResponseBody) *UpdateExperimentFolderResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentFolderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
