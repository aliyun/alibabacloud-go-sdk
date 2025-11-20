// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContainerClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContainerClusterResponse
	GetStatusCode() *int32
	SetBody(v *AddContainerClusterResponseBody) *AddContainerClusterResponse
	GetBody() *AddContainerClusterResponseBody
}

type AddContainerClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContainerClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *AddContainerClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContainerClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContainerClusterResponse) GetBody() *AddContainerClusterResponseBody {
	return s.Body
}

func (s *AddContainerClusterResponse) SetHeaders(v map[string]*string) *AddContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *AddContainerClusterResponse) SetStatusCode(v int32) *AddContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContainerClusterResponse) SetBody(v *AddContainerClusterResponseBody) *AddContainerClusterResponse {
	s.Body = v
	return s
}

func (s *AddContainerClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
