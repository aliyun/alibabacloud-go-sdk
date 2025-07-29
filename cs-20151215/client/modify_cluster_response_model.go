// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterResponseBody) *ModifyClusterResponse
	GetBody() *ModifyClusterResponseBody
}

type ModifyClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterResponse) GetBody() *ModifyClusterResponseBody {
	return s.Body
}

func (s *ModifyClusterResponse) SetHeaders(v map[string]*string) *ModifyClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterResponse) SetStatusCode(v int32) *ModifyClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterResponse) SetBody(v *ModifyClusterResponseBody) *ModifyClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterResponse) Validate() error {
	return dara.Validate(s)
}
