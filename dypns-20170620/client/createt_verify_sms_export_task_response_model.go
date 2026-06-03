// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatetVerifySmsExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatetVerifySmsExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatetVerifySmsExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatetVerifySmsExportTaskResponseBody) *CreatetVerifySmsExportTaskResponse
	GetBody() *CreatetVerifySmsExportTaskResponseBody
}

type CreatetVerifySmsExportTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatetVerifySmsExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatetVerifySmsExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatetVerifySmsExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatetVerifySmsExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatetVerifySmsExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatetVerifySmsExportTaskResponse) GetBody() *CreatetVerifySmsExportTaskResponseBody {
	return s.Body
}

func (s *CreatetVerifySmsExportTaskResponse) SetHeaders(v map[string]*string) *CreatetVerifySmsExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatetVerifySmsExportTaskResponse) SetStatusCode(v int32) *CreatetVerifySmsExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponse) SetBody(v *CreatetVerifySmsExportTaskResponseBody) *CreatetVerifySmsExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreatetVerifySmsExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
