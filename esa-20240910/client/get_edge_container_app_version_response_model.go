// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppVersionResponseBody) *GetEdgeContainerAppVersionResponse
	GetBody() *GetEdgeContainerAppVersionResponseBody
}

type GetEdgeContainerAppVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppVersionResponse) GetBody() *GetEdgeContainerAppVersionResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppVersionResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppVersionResponse) SetStatusCode(v int32) *GetEdgeContainerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponse) SetBody(v *GetEdgeContainerAppVersionResponseBody) *GetEdgeContainerAppVersionResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
