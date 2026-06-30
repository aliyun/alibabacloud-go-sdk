// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeProjectExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeProjectExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeProjectExportJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeProjectExportJobResponseBody) *SubmitYikeProjectExportJobResponse
	GetBody() *SubmitYikeProjectExportJobResponseBody
}

type SubmitYikeProjectExportJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeProjectExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeProjectExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeProjectExportJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeProjectExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeProjectExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeProjectExportJobResponse) GetBody() *SubmitYikeProjectExportJobResponseBody {
	return s.Body
}

func (s *SubmitYikeProjectExportJobResponse) SetHeaders(v map[string]*string) *SubmitYikeProjectExportJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeProjectExportJobResponse) SetStatusCode(v int32) *SubmitYikeProjectExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeProjectExportJobResponse) SetBody(v *SubmitYikeProjectExportJobResponseBody) *SubmitYikeProjectExportJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeProjectExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
