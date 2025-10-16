// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEcdReportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEcdReportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEcdReportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateEcdReportTaskResponseBody) *CreateEcdReportTaskResponse
	GetBody() *CreateEcdReportTaskResponseBody
}

type CreateEcdReportTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEcdReportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEcdReportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEcdReportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateEcdReportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEcdReportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEcdReportTaskResponse) GetBody() *CreateEcdReportTaskResponseBody {
	return s.Body
}

func (s *CreateEcdReportTaskResponse) SetHeaders(v map[string]*string) *CreateEcdReportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateEcdReportTaskResponse) SetStatusCode(v int32) *CreateEcdReportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEcdReportTaskResponse) SetBody(v *CreateEcdReportTaskResponseBody) *CreateEcdReportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateEcdReportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
