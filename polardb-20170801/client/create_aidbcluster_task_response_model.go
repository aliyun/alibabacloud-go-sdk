// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIDBClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIDBClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIDBClusterTaskResponseBody) *CreateAIDBClusterTaskResponse
	GetBody() *CreateAIDBClusterTaskResponseBody
}

type CreateAIDBClusterTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIDBClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIDBClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIDBClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIDBClusterTaskResponse) GetBody() *CreateAIDBClusterTaskResponseBody {
	return s.Body
}

func (s *CreateAIDBClusterTaskResponse) SetHeaders(v map[string]*string) *CreateAIDBClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAIDBClusterTaskResponse) SetStatusCode(v int32) *CreateAIDBClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIDBClusterTaskResponse) SetBody(v *CreateAIDBClusterTaskResponseBody) *CreateAIDBClusterTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAIDBClusterTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
