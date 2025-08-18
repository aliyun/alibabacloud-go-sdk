// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEdgeContainerAppVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackEdgeContainerAppVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackEdgeContainerAppVersionResponse
	GetStatusCode() *int32
	SetBody(v *RollbackEdgeContainerAppVersionResponseBody) *RollbackEdgeContainerAppVersionResponse
	GetBody() *RollbackEdgeContainerAppVersionResponseBody
}

type RollbackEdgeContainerAppVersionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackEdgeContainerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackEdgeContainerAppVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackEdgeContainerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackEdgeContainerAppVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackEdgeContainerAppVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackEdgeContainerAppVersionResponse) GetBody() *RollbackEdgeContainerAppVersionResponseBody {
	return s.Body
}

func (s *RollbackEdgeContainerAppVersionResponse) SetHeaders(v map[string]*string) *RollbackEdgeContainerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackEdgeContainerAppVersionResponse) SetStatusCode(v int32) *RollbackEdgeContainerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionResponse) SetBody(v *RollbackEdgeContainerAppVersionResponseBody) *RollbackEdgeContainerAppVersionResponse {
	s.Body = v
	return s
}

func (s *RollbackEdgeContainerAppVersionResponse) Validate() error {
	return dara.Validate(s)
}
