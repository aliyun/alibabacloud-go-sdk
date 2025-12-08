// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVirtualClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVirtualClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVirtualClusterResponseBody) *ModifyVirtualClusterResponse
	GetBody() *ModifyVirtualClusterResponseBody
}

type ModifyVirtualClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVirtualClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVirtualClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyVirtualClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVirtualClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVirtualClusterResponse) GetBody() *ModifyVirtualClusterResponseBody {
	return s.Body
}

func (s *ModifyVirtualClusterResponse) SetHeaders(v map[string]*string) *ModifyVirtualClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyVirtualClusterResponse) SetStatusCode(v int32) *ModifyVirtualClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVirtualClusterResponse) SetBody(v *ModifyVirtualClusterResponseBody) *ModifyVirtualClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyVirtualClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
