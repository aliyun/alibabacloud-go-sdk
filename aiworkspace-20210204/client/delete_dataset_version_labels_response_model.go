// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetVersionLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetVersionLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetVersionLabelsResponseBody) *DeleteDatasetVersionLabelsResponse
	GetBody() *DeleteDatasetVersionLabelsResponseBody
}

type DeleteDatasetVersionLabelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetVersionLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetVersionLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetVersionLabelsResponse) GetBody() *DeleteDatasetVersionLabelsResponseBody {
	return s.Body
}

func (s *DeleteDatasetVersionLabelsResponse) SetHeaders(v map[string]*string) *DeleteDatasetVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetVersionLabelsResponse) SetStatusCode(v int32) *DeleteDatasetVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetVersionLabelsResponse) SetBody(v *DeleteDatasetVersionLabelsResponseBody) *DeleteDatasetVersionLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetVersionLabelsResponse) Validate() error {
	return dara.Validate(s)
}
