// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFabricTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFabricTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFabricTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetFabricTopologyResponseBody) *GetFabricTopologyResponse
	GetBody() *GetFabricTopologyResponseBody
}

type GetFabricTopologyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFabricTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFabricTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFabricTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFabricTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFabricTopologyResponse) GetBody() *GetFabricTopologyResponseBody {
	return s.Body
}

func (s *GetFabricTopologyResponse) SetHeaders(v map[string]*string) *GetFabricTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetFabricTopologyResponse) SetStatusCode(v int32) *GetFabricTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFabricTopologyResponse) SetBody(v *GetFabricTopologyResponseBody) *GetFabricTopologyResponse {
	s.Body = v
	return s
}

func (s *GetFabricTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
