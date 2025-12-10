// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentFolderResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentFolderResponseBody) *CreateExperimentFolderResponse
	GetBody() *CreateExperimentFolderResponseBody
}

type CreateExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentFolderResponse) GetBody() *CreateExperimentFolderResponseBody {
	return s.Body
}

func (s *CreateExperimentFolderResponse) SetHeaders(v map[string]*string) *CreateExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentFolderResponse) SetStatusCode(v int32) *CreateExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentFolderResponse) SetBody(v *CreateExperimentFolderResponseBody) *CreateExperimentFolderResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentFolderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
