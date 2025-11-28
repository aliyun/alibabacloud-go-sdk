// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricConsortiumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSFabricConsortiumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSFabricConsortiumResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSFabricConsortiumResponseBody) *ListBaaSFabricConsortiumResponse
	GetBody() *ListBaaSFabricConsortiumResponseBody
}

type ListBaaSFabricConsortiumResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSFabricConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSFabricConsortiumResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricConsortiumResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricConsortiumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSFabricConsortiumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSFabricConsortiumResponse) GetBody() *ListBaaSFabricConsortiumResponseBody {
	return s.Body
}

func (s *ListBaaSFabricConsortiumResponse) SetHeaders(v map[string]*string) *ListBaaSFabricConsortiumResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSFabricConsortiumResponse) SetStatusCode(v int32) *ListBaaSFabricConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponse) SetBody(v *ListBaaSFabricConsortiumResponseBody) *ListBaaSFabricConsortiumResponse {
	s.Body = v
	return s
}

func (s *ListBaaSFabricConsortiumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
