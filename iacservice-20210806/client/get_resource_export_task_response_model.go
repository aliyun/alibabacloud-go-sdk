// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceExportTaskResponseBody) *GetResourceExportTaskResponse
	GetBody() *GetResourceExportTaskResponseBody
}

type GetResourceExportTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceExportTaskResponse) GetBody() *GetResourceExportTaskResponseBody {
	return s.Body
}

func (s *GetResourceExportTaskResponse) SetHeaders(v map[string]*string) *GetResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetResourceExportTaskResponse) SetStatusCode(v int32) *GetResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceExportTaskResponse) SetBody(v *GetResourceExportTaskResponseBody) *GetResourceExportTaskResponse {
	s.Body = v
	return s
}

func (s *GetResourceExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
