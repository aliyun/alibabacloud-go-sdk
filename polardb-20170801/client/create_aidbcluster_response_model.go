// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIDBClusterResponseBody) *CreateAIDBClusterResponse
	GetBody() *CreateAIDBClusterResponseBody
}

type CreateAIDBClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIDBClusterResponse) GetBody() *CreateAIDBClusterResponseBody {
	return s.Body
}

func (s *CreateAIDBClusterResponse) SetHeaders(v map[string]*string) *CreateAIDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateAIDBClusterResponse) SetStatusCode(v int32) *CreateAIDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIDBClusterResponse) SetBody(v *CreateAIDBClusterResponseBody) *CreateAIDBClusterResponse {
	s.Body = v
	return s
}

func (s *CreateAIDBClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
