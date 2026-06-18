// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerResponseBody) *GetEdgeContainerResponse
	GetBody() *GetEdgeContainerResponseBody
}

type GetEdgeContainerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerResponse) GetBody() *GetEdgeContainerResponseBody {
	return s.Body
}

func (s *GetEdgeContainerResponse) SetHeaders(v map[string]*string) *GetEdgeContainerResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerResponse) SetStatusCode(v int32) *GetEdgeContainerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerResponse) SetBody(v *GetEdgeContainerResponseBody) *GetEdgeContainerResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
