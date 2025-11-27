// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClusteringTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFigureClusteringTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFigureClusteringTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFigureClusteringTaskResponseBody) *CreateFigureClusteringTaskResponse
	GetBody() *CreateFigureClusteringTaskResponseBody
}

type CreateFigureClusteringTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFigureClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFigureClusteringTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFigureClusteringTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFigureClusteringTaskResponse) GetBody() *CreateFigureClusteringTaskResponseBody {
	return s.Body
}

func (s *CreateFigureClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateFigureClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFigureClusteringTaskResponse) SetStatusCode(v int32) *CreateFigureClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFigureClusteringTaskResponse) SetBody(v *CreateFigureClusteringTaskResponseBody) *CreateFigureClusteringTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFigureClusteringTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
