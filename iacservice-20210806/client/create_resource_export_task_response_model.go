// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceExportTaskResponseBody) *CreateResourceExportTaskResponse
	GetBody() *CreateResourceExportTaskResponseBody
}

type CreateResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceExportTaskResponse) GetBody() *CreateResourceExportTaskResponseBody {
	return s.Body
}

func (s *CreateResourceExportTaskResponse) SetHeaders(v map[string]*string) *CreateResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceExportTaskResponse) SetStatusCode(v int32) *CreateResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceExportTaskResponse) SetBody(v *CreateResourceExportTaskResponseBody) *CreateResourceExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateResourceExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
