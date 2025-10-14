// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppStatusResponseBody) *GetEdgeContainerAppStatusResponse
	GetBody() *GetEdgeContainerAppStatusResponseBody
}

type GetEdgeContainerAppStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppStatusResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppStatusResponse) GetBody() *GetEdgeContainerAppStatusResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppStatusResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppStatusResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppStatusResponse) SetStatusCode(v int32) *GetEdgeContainerAppStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponse) SetBody(v *GetEdgeContainerAppStatusResponseBody) *GetEdgeContainerAppStatusResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
