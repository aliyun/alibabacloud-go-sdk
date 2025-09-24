// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetJobResponseBody) *DeleteDatasetJobResponse
	GetBody() *DeleteDatasetJobResponseBody
}

type DeleteDatasetJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetJobResponse) GetBody() *DeleteDatasetJobResponseBody {
	return s.Body
}

func (s *DeleteDatasetJobResponse) SetHeaders(v map[string]*string) *DeleteDatasetJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetJobResponse) SetStatusCode(v int32) *DeleteDatasetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetJobResponse) SetBody(v *DeleteDatasetJobResponseBody) *DeleteDatasetJobResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetJobResponse) Validate() error {
	return dara.Validate(s)
}
