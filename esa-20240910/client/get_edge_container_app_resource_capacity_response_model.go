// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppResourceCapacityResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppResourceCapacityResponseBody) *GetEdgeContainerAppResourceCapacityResponse
	GetBody() *GetEdgeContainerAppResourceCapacityResponseBody
}

type GetEdgeContainerAppResourceCapacityResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppResourceCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppResourceCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppResourceCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppResourceCapacityResponse) GetBody() *GetEdgeContainerAppResourceCapacityResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppResourceCapacityResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponse) SetStatusCode(v int32) *GetEdgeContainerAppResourceCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponse) SetBody(v *GetEdgeContainerAppResourceCapacityResponseBody) *GetEdgeContainerAppResourceCapacityResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
