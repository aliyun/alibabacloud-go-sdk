// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportHttpApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchExportHttpApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchExportHttpApisResponse
	GetStatusCode() *int32
	SetBody(v *BatchExportHttpApisResponseBody) *BatchExportHttpApisResponse
	GetBody() *BatchExportHttpApisResponseBody
}

type BatchExportHttpApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchExportHttpApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchExportHttpApisResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchExportHttpApisResponse) GoString() string {
	return s.String()
}

func (s *BatchExportHttpApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchExportHttpApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchExportHttpApisResponse) GetBody() *BatchExportHttpApisResponseBody {
	return s.Body
}

func (s *BatchExportHttpApisResponse) SetHeaders(v map[string]*string) *BatchExportHttpApisResponse {
	s.Headers = v
	return s
}

func (s *BatchExportHttpApisResponse) SetStatusCode(v int32) *BatchExportHttpApisResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExportHttpApisResponse) SetBody(v *BatchExportHttpApisResponseBody) *BatchExportHttpApisResponse {
	s.Body = v
	return s
}

func (s *BatchExportHttpApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
