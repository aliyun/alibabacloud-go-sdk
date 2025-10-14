// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeContainerAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeContainerAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeContainerAppResponseBody) *CreateEdgeContainerAppResponse
	GetBody() *CreateEdgeContainerAppResponseBody
}

type CreateEdgeContainerAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeContainerAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeContainerAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeContainerAppResponse) GetBody() *CreateEdgeContainerAppResponseBody {
	return s.Body
}

func (s *CreateEdgeContainerAppResponse) SetHeaders(v map[string]*string) *CreateEdgeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeContainerAppResponse) SetStatusCode(v int32) *CreateEdgeContainerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeContainerAppResponse) SetBody(v *CreateEdgeContainerAppResponseBody) *CreateEdgeContainerAppResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeContainerAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
