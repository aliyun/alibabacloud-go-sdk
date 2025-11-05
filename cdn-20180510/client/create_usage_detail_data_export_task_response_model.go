// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsageDetailDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUsageDetailDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUsageDetailDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateUsageDetailDataExportTaskResponseBody) *CreateUsageDetailDataExportTaskResponse
	GetBody() *CreateUsageDetailDataExportTaskResponseBody
}

type CreateUsageDetailDataExportTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUsageDetailDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUsageDetailDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUsageDetailDataExportTaskResponse) GetBody() *CreateUsageDetailDataExportTaskResponseBody {
	return s.Body
}

func (s *CreateUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *CreateUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponse) SetStatusCode(v int32) *CreateUsageDetailDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponse) SetBody(v *CreateUsageDetailDataExportTaskResponseBody) *CreateUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
