// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetVersionLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetVersionLabelsResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetVersionLabelsResponseBody) *CreateDatasetVersionLabelsResponse
	GetBody() *CreateDatasetVersionLabelsResponseBody
}

type CreateDatasetVersionLabelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetVersionLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetVersionLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetVersionLabelsResponse) GetBody() *CreateDatasetVersionLabelsResponseBody {
	return s.Body
}

func (s *CreateDatasetVersionLabelsResponse) SetHeaders(v map[string]*string) *CreateDatasetVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetVersionLabelsResponse) SetStatusCode(v int32) *CreateDatasetVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetVersionLabelsResponse) SetBody(v *CreateDatasetVersionLabelsResponseBody) *CreateDatasetVersionLabelsResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetVersionLabelsResponse) Validate() error {
	return dara.Validate(s)
}
