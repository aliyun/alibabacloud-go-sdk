// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridMonitorNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridMonitorNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridMonitorNamespaceResponseBody) *DeleteHybridMonitorNamespaceResponse
	GetBody() *DeleteHybridMonitorNamespaceResponseBody
}

type DeleteHybridMonitorNamespaceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridMonitorNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridMonitorNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridMonitorNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridMonitorNamespaceResponse) GetBody() *DeleteHybridMonitorNamespaceResponseBody {
	return s.Body
}

func (s *DeleteHybridMonitorNamespaceResponse) SetHeaders(v map[string]*string) *DeleteHybridMonitorNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponse) SetStatusCode(v int32) *DeleteHybridMonitorNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponse) SetBody(v *DeleteHybridMonitorNamespaceResponseBody) *DeleteHybridMonitorNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
