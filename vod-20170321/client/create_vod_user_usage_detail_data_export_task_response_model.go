// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodUserUsageDetailDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVodUserUsageDetailDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVodUserUsageDetailDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVodUserUsageDetailDataExportTaskResponseBody) *CreateVodUserUsageDetailDataExportTaskResponse
	GetBody() *CreateVodUserUsageDetailDataExportTaskResponseBody
}

type CreateVodUserUsageDetailDataExportTaskResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVodUserUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVodUserUsageDetailDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVodUserUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) GetBody() *CreateVodUserUsageDetailDataExportTaskResponseBody {
	return s.Body
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *CreateVodUserUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) SetStatusCode(v int32) *CreateVodUserUsageDetailDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) SetBody(v *CreateVodUserUsageDetailDataExportTaskResponseBody) *CreateVodUserUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
