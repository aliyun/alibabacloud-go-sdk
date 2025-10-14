// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchExportResponse
	GetStatusCode() *int32
	SetBody(v *BatchExportResponseBody) *BatchExportResponse
	GetBody() *BatchExportResponseBody
}

type BatchExportResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchExportResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchExportResponse) GoString() string {
	return s.String()
}

func (s *BatchExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchExportResponse) GetBody() *BatchExportResponseBody {
	return s.Body
}

func (s *BatchExportResponse) SetHeaders(v map[string]*string) *BatchExportResponse {
	s.Headers = v
	return s
}

func (s *BatchExportResponse) SetStatusCode(v int32) *BatchExportResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExportResponse) SetBody(v *BatchExportResponseBody) *BatchExportResponse {
	s.Body = v
	return s
}

func (s *BatchExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
