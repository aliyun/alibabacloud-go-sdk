// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeContainerAppVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeContainerAppVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeContainerAppVersionResponseBody) *CreateEdgeContainerAppVersionResponse
	GetBody() *CreateEdgeContainerAppVersionResponseBody
}

type CreateEdgeContainerAppVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeContainerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeContainerAppVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeContainerAppVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeContainerAppVersionResponse) GetBody() *CreateEdgeContainerAppVersionResponseBody {
	return s.Body
}

func (s *CreateEdgeContainerAppVersionResponse) SetHeaders(v map[string]*string) *CreateEdgeContainerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeContainerAppVersionResponse) SetStatusCode(v int32) *CreateEdgeContainerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeContainerAppVersionResponse) SetBody(v *CreateEdgeContainerAppVersionResponseBody) *CreateEdgeContainerAppVersionResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeContainerAppVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
