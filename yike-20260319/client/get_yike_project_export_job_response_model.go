// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeProjectExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeProjectExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeProjectExportJobResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeProjectExportJobResponseBody) *GetYikeProjectExportJobResponse
	GetBody() *GetYikeProjectExportJobResponseBody
}

type GetYikeProjectExportJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeProjectExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeProjectExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobResponse) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeProjectExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeProjectExportJobResponse) GetBody() *GetYikeProjectExportJobResponseBody {
	return s.Body
}

func (s *GetYikeProjectExportJobResponse) SetHeaders(v map[string]*string) *GetYikeProjectExportJobResponse {
	s.Headers = v
	return s
}

func (s *GetYikeProjectExportJobResponse) SetStatusCode(v int32) *GetYikeProjectExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeProjectExportJobResponse) SetBody(v *GetYikeProjectExportJobResponseBody) *GetYikeProjectExportJobResponse {
	s.Body = v
	return s
}

func (s *GetYikeProjectExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
