// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAIDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartAIDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartAIDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *RestartAIDBClusterResponseBody) *RestartAIDBClusterResponse
	GetBody() *RestartAIDBClusterResponseBody
}

type RestartAIDBClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartAIDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartAIDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartAIDBClusterResponse) GoString() string {
	return s.String()
}

func (s *RestartAIDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartAIDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartAIDBClusterResponse) GetBody() *RestartAIDBClusterResponseBody {
	return s.Body
}

func (s *RestartAIDBClusterResponse) SetHeaders(v map[string]*string) *RestartAIDBClusterResponse {
	s.Headers = v
	return s
}

func (s *RestartAIDBClusterResponse) SetStatusCode(v int32) *RestartAIDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartAIDBClusterResponse) SetBody(v *RestartAIDBClusterResponseBody) *RestartAIDBClusterResponse {
	s.Body = v
	return s
}

func (s *RestartAIDBClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
