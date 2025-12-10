// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentFolderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentFolderResponseBody) *DeleteExperimentFolderResponse
	GetBody() *DeleteExperimentFolderResponseBody
}

type DeleteExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentFolderResponse) GetBody() *DeleteExperimentFolderResponseBody {
	return s.Body
}

func (s *DeleteExperimentFolderResponse) SetHeaders(v map[string]*string) *DeleteExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentFolderResponse) SetStatusCode(v int32) *DeleteExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentFolderResponse) SetBody(v *DeleteExperimentFolderResponseBody) *DeleteExperimentFolderResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentFolderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
