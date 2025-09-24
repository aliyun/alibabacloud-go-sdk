// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetLabelsResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetLabelsResponseBody) *CreateDatasetLabelsResponse
	GetBody() *CreateDatasetLabelsResponseBody
}

type CreateDatasetLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetLabelsResponse) GetBody() *CreateDatasetLabelsResponseBody {
	return s.Body
}

func (s *CreateDatasetLabelsResponse) SetHeaders(v map[string]*string) *CreateDatasetLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetLabelsResponse) SetStatusCode(v int32) *CreateDatasetLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetLabelsResponse) SetBody(v *CreateDatasetLabelsResponseBody) *CreateDatasetLabelsResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetLabelsResponse) Validate() error {
	return dara.Validate(s)
}
