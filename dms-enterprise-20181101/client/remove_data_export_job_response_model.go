// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataExportJobResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDataExportJobResponseBody) *RemoveDataExportJobResponse
	GetBody() *RemoveDataExportJobResponseBody
}

type RemoveDataExportJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataExportJobResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataExportJobResponse) GetBody() *RemoveDataExportJobResponseBody {
	return s.Body
}

func (s *RemoveDataExportJobResponse) SetHeaders(v map[string]*string) *RemoveDataExportJobResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataExportJobResponse) SetStatusCode(v int32) *RemoveDataExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataExportJobResponse) SetBody(v *RemoveDataExportJobResponseBody) *RemoveDataExportJobResponse {
	s.Body = v
	return s
}

func (s *RemoveDataExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
