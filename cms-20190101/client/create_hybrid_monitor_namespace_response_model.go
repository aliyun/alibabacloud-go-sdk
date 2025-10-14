// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridMonitorNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridMonitorNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridMonitorNamespaceResponseBody) *CreateHybridMonitorNamespaceResponse
	GetBody() *CreateHybridMonitorNamespaceResponseBody
}

type CreateHybridMonitorNamespaceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridMonitorNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridMonitorNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridMonitorNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridMonitorNamespaceResponse) GetBody() *CreateHybridMonitorNamespaceResponseBody {
	return s.Body
}

func (s *CreateHybridMonitorNamespaceResponse) SetHeaders(v map[string]*string) *CreateHybridMonitorNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridMonitorNamespaceResponse) SetStatusCode(v int32) *CreateHybridMonitorNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridMonitorNamespaceResponse) SetBody(v *CreateHybridMonitorNamespaceResponseBody) *CreateHybridMonitorNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateHybridMonitorNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
