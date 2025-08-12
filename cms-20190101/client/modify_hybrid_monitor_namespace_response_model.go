// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridMonitorNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridMonitorNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridMonitorNamespaceResponseBody) *ModifyHybridMonitorNamespaceResponse
	GetBody() *ModifyHybridMonitorNamespaceResponseBody
}

type ModifyHybridMonitorNamespaceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridMonitorNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridMonitorNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridMonitorNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridMonitorNamespaceResponse) GetBody() *ModifyHybridMonitorNamespaceResponseBody {
	return s.Body
}

func (s *ModifyHybridMonitorNamespaceResponse) SetHeaders(v map[string]*string) *ModifyHybridMonitorNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponse) SetStatusCode(v int32) *ModifyHybridMonitorNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponse) SetBody(v *ModifyHybridMonitorNamespaceResponseBody) *ModifyHybridMonitorNamespaceResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
