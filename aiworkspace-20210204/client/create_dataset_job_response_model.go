// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetJobResponseBody) *CreateDatasetJobResponse
	GetBody() *CreateDatasetJobResponseBody
}

type CreateDatasetJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetJobResponse) GetBody() *CreateDatasetJobResponseBody {
	return s.Body
}

func (s *CreateDatasetJobResponse) SetHeaders(v map[string]*string) *CreateDatasetJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetJobResponse) SetStatusCode(v int32) *CreateDatasetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetJobResponse) SetBody(v *CreateDatasetJobResponseBody) *CreateDatasetJobResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetJobResponse) Validate() error {
	return dara.Validate(s)
}
