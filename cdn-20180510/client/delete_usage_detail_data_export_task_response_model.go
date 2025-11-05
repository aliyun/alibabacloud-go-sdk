// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsageDetailDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUsageDetailDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUsageDetailDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUsageDetailDataExportTaskResponseBody) *DeleteUsageDetailDataExportTaskResponse
	GetBody() *DeleteUsageDetailDataExportTaskResponseBody
}

type DeleteUsageDetailDataExportTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUsageDetailDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUsageDetailDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUsageDetailDataExportTaskResponse) GetBody() *DeleteUsageDetailDataExportTaskResponseBody {
	return s.Body
}

func (s *DeleteUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *DeleteUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsageDetailDataExportTaskResponse) SetStatusCode(v int32) *DeleteUsageDetailDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUsageDetailDataExportTaskResponse) SetBody(v *DeleteUsageDetailDataExportTaskResponseBody) *DeleteUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteUsageDetailDataExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
