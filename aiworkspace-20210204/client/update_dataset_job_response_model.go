// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetJobResponseBody) *UpdateDatasetJobResponse
	GetBody() *UpdateDatasetJobResponseBody
}

type UpdateDatasetJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetJobResponse) GetBody() *UpdateDatasetJobResponseBody {
	return s.Body
}

func (s *UpdateDatasetJobResponse) SetHeaders(v map[string]*string) *UpdateDatasetJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetJobResponse) SetStatusCode(v int32) *UpdateDatasetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetJobResponse) SetBody(v *UpdateDatasetJobResponseBody) *UpdateDatasetJobResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetJobResponse) Validate() error {
	return dara.Validate(s)
}
