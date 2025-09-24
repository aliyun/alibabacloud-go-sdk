// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetLabelsResponseBody) *DeleteDatasetLabelsResponse
	GetBody() *DeleteDatasetLabelsResponseBody
}

type DeleteDatasetLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetLabelsResponse) GetBody() *DeleteDatasetLabelsResponseBody {
	return s.Body
}

func (s *DeleteDatasetLabelsResponse) SetHeaders(v map[string]*string) *DeleteDatasetLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetLabelsResponse) SetStatusCode(v int32) *DeleteDatasetLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetLabelsResponse) SetBody(v *DeleteDatasetLabelsResponseBody) *DeleteDatasetLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetLabelsResponse) Validate() error {
	return dara.Validate(s)
}
