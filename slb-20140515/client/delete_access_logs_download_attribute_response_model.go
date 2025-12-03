// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessLogsDownloadAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessLogsDownloadAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessLogsDownloadAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessLogsDownloadAttributeResponseBody) *DeleteAccessLogsDownloadAttributeResponse
	GetBody() *DeleteAccessLogsDownloadAttributeResponseBody
}

type DeleteAccessLogsDownloadAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessLogsDownloadAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessLogsDownloadAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessLogsDownloadAttributeResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessLogsDownloadAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessLogsDownloadAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessLogsDownloadAttributeResponse) GetBody() *DeleteAccessLogsDownloadAttributeResponseBody {
	return s.Body
}

func (s *DeleteAccessLogsDownloadAttributeResponse) SetHeaders(v map[string]*string) *DeleteAccessLogsDownloadAttributeResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeResponse) SetStatusCode(v int32) *DeleteAccessLogsDownloadAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeResponse) SetBody(v *DeleteAccessLogsDownloadAttributeResponseBody) *DeleteAccessLogsDownloadAttributeResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
