// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudImportTaskTelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudImportTaskTelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudImportTaskTelResponse
	GetStatusCode() *int32
	SetBody(v *CloudImportTaskTelResponseBody) *CloudImportTaskTelResponse
	GetBody() *CloudImportTaskTelResponseBody
}

type CloudImportTaskTelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudImportTaskTelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudImportTaskTelResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelResponse) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudImportTaskTelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudImportTaskTelResponse) GetBody() *CloudImportTaskTelResponseBody {
	return s.Body
}

func (s *CloudImportTaskTelResponse) SetHeaders(v map[string]*string) *CloudImportTaskTelResponse {
	s.Headers = v
	return s
}

func (s *CloudImportTaskTelResponse) SetStatusCode(v int32) *CloudImportTaskTelResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudImportTaskTelResponse) SetBody(v *CloudImportTaskTelResponseBody) *CloudImportTaskTelResponse {
	s.Body = v
	return s
}

func (s *CloudImportTaskTelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
