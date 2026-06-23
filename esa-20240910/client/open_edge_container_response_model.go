// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEdgeContainerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenEdgeContainerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenEdgeContainerResponse
	GetStatusCode() *int32
	SetBody(v *OpenEdgeContainerResponseBody) *OpenEdgeContainerResponse
	GetBody() *OpenEdgeContainerResponseBody
}

type OpenEdgeContainerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenEdgeContainerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenEdgeContainerResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenEdgeContainerResponse) GoString() string {
	return s.String()
}

func (s *OpenEdgeContainerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenEdgeContainerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenEdgeContainerResponse) GetBody() *OpenEdgeContainerResponseBody {
	return s.Body
}

func (s *OpenEdgeContainerResponse) SetHeaders(v map[string]*string) *OpenEdgeContainerResponse {
	s.Headers = v
	return s
}

func (s *OpenEdgeContainerResponse) SetStatusCode(v int32) *OpenEdgeContainerResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenEdgeContainerResponse) SetBody(v *OpenEdgeContainerResponseBody) *OpenEdgeContainerResponse {
	s.Body = v
	return s
}

func (s *OpenEdgeContainerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
