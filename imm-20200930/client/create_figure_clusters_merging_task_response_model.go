// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClustersMergingTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFigureClustersMergingTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFigureClustersMergingTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFigureClustersMergingTaskResponseBody) *CreateFigureClustersMergingTaskResponse
	GetBody() *CreateFigureClustersMergingTaskResponseBody
}

type CreateFigureClustersMergingTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFigureClustersMergingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFigureClustersMergingTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClustersMergingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFigureClustersMergingTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFigureClustersMergingTaskResponse) GetBody() *CreateFigureClustersMergingTaskResponseBody {
	return s.Body
}

func (s *CreateFigureClustersMergingTaskResponse) SetHeaders(v map[string]*string) *CreateFigureClustersMergingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFigureClustersMergingTaskResponse) SetStatusCode(v int32) *CreateFigureClustersMergingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponse) SetBody(v *CreateFigureClustersMergingTaskResponseBody) *CreateFigureClustersMergingTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFigureClustersMergingTaskResponse) Validate() error {
	return dara.Validate(s)
}
