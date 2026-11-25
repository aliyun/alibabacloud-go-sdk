// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEditMediaManifestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitEditMediaManifestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitEditMediaManifestJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitEditMediaManifestJobResponseBody) *SubmitEditMediaManifestJobResponse
	GetBody() *SubmitEditMediaManifestJobResponseBody
}

type SubmitEditMediaManifestJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEditMediaManifestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEditMediaManifestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitEditMediaManifestJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitEditMediaManifestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitEditMediaManifestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitEditMediaManifestJobResponse) GetBody() *SubmitEditMediaManifestJobResponseBody {
	return s.Body
}

func (s *SubmitEditMediaManifestJobResponse) SetHeaders(v map[string]*string) *SubmitEditMediaManifestJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitEditMediaManifestJobResponse) SetStatusCode(v int32) *SubmitEditMediaManifestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEditMediaManifestJobResponse) SetBody(v *SubmitEditMediaManifestJobResponseBody) *SubmitEditMediaManifestJobResponse {
	s.Body = v
	return s
}

func (s *SubmitEditMediaManifestJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
