// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppResponseBody) *GetEdgeContainerAppResponse
	GetBody() *GetEdgeContainerAppResponseBody
}

type GetEdgeContainerAppResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppResponse) GetBody() *GetEdgeContainerAppResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppResponse) SetStatusCode(v int32) *GetEdgeContainerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppResponse) SetBody(v *GetEdgeContainerAppResponseBody) *GetEdgeContainerAppResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
