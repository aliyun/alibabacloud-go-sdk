// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppResourceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppResourceStatusResponseBody) *GetEdgeContainerAppResourceStatusResponse
	GetBody() *GetEdgeContainerAppResourceStatusResponseBody
}

type GetEdgeContainerAppResourceStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppResourceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppResourceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppResourceStatusResponse) GetBody() *GetEdgeContainerAppResourceStatusResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppResourceStatusResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponse) SetStatusCode(v int32) *GetEdgeContainerAppResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponse) SetBody(v *GetEdgeContainerAppResourceStatusResponseBody) *GetEdgeContainerAppResourceStatusResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
