// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExplorerModuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExplorerModuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExplorerModuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExplorerModuleAttributeResponseBody) *UpdateExplorerModuleAttributeResponse
	GetBody() *UpdateExplorerModuleAttributeResponseBody
}

type UpdateExplorerModuleAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExplorerModuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExplorerModuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExplorerModuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateExplorerModuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExplorerModuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExplorerModuleAttributeResponse) GetBody() *UpdateExplorerModuleAttributeResponseBody {
	return s.Body
}

func (s *UpdateExplorerModuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateExplorerModuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateExplorerModuleAttributeResponse) SetStatusCode(v int32) *UpdateExplorerModuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExplorerModuleAttributeResponse) SetBody(v *UpdateExplorerModuleAttributeResponseBody) *UpdateExplorerModuleAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateExplorerModuleAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
