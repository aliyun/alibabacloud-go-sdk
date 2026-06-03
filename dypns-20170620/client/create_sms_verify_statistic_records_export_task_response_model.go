// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsVerifyStatisticRecordsExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsVerifyStatisticRecordsExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsVerifyStatisticRecordsExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) *CreateSmsVerifyStatisticRecordsExportTaskResponse
	GetBody() *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
}

type CreateSmsVerifyStatisticRecordsExportTaskResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsVerifyStatisticRecordsExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsVerifyStatisticRecordsExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsVerifyStatisticRecordsExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) GetBody() *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	return s.Body
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) SetHeaders(v map[string]*string) *CreateSmsVerifyStatisticRecordsExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) SetStatusCode(v int32) *CreateSmsVerifyStatisticRecordsExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) SetBody(v *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) *CreateSmsVerifyStatisticRecordsExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
