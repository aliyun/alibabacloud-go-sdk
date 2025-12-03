// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessLogsDownloadAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccessLogsDownloadAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccessLogsDownloadAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetAccessLogsDownloadAttributeResponseBody) *SetAccessLogsDownloadAttributeResponse
	GetBody() *SetAccessLogsDownloadAttributeResponseBody
}

type SetAccessLogsDownloadAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessLogsDownloadAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessLogsDownloadAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccessLogsDownloadAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetAccessLogsDownloadAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccessLogsDownloadAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccessLogsDownloadAttributeResponse) GetBody() *SetAccessLogsDownloadAttributeResponseBody {
	return s.Body
}

func (s *SetAccessLogsDownloadAttributeResponse) SetHeaders(v map[string]*string) *SetAccessLogsDownloadAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetAccessLogsDownloadAttributeResponse) SetStatusCode(v int32) *SetAccessLogsDownloadAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeResponse) SetBody(v *SetAccessLogsDownloadAttributeResponseBody) *SetAccessLogsDownloadAttributeResponse {
	s.Body = v
	return s
}

func (s *SetAccessLogsDownloadAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
