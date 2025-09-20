// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationDateClusteringTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLocationDateClusteringTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLocationDateClusteringTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateLocationDateClusteringTaskResponseBody) *CreateLocationDateClusteringTaskResponse
	GetBody() *CreateLocationDateClusteringTaskResponseBody
}

type CreateLocationDateClusteringTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocationDateClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocationDateClusteringTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLocationDateClusteringTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLocationDateClusteringTaskResponse) GetBody() *CreateLocationDateClusteringTaskResponseBody {
	return s.Body
}

func (s *CreateLocationDateClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateLocationDateClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLocationDateClusteringTaskResponse) SetStatusCode(v int32) *CreateLocationDateClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponse) SetBody(v *CreateLocationDateClusteringTaskResponseBody) *CreateLocationDateClusteringTaskResponse {
	s.Body = v
	return s
}

func (s *CreateLocationDateClusteringTaskResponse) Validate() error {
	return dara.Validate(s)
}
