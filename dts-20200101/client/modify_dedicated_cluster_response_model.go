// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedClusterResponseBody) *ModifyDedicatedClusterResponse
	GetBody() *ModifyDedicatedClusterResponseBody
}

type ModifyDedicatedClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedClusterResponse) GetBody() *ModifyDedicatedClusterResponseBody {
	return s.Body
}

func (s *ModifyDedicatedClusterResponse) SetHeaders(v map[string]*string) *ModifyDedicatedClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedClusterResponse) SetStatusCode(v int32) *ModifyDedicatedClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedClusterResponse) SetBody(v *ModifyDedicatedClusterResponseBody) *ModifyDedicatedClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedClusterResponse) Validate() error {
	return dara.Validate(s)
}
