// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse
	GetBody() *ModifyDBClusterResponseBody
}

type ModifyDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterResponse) GetBody() *ModifyDBClusterResponseBody {
	return s.Body
}

func (s *ModifyDBClusterResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResponse) SetStatusCode(v int32) *ModifyDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
