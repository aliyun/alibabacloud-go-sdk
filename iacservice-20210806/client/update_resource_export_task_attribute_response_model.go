// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceExportTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceExportTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceExportTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceExportTaskAttributeResponseBody) *UpdateResourceExportTaskAttributeResponse
	GetBody() *UpdateResourceExportTaskAttributeResponseBody
}

type UpdateResourceExportTaskAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceExportTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceExportTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceExportTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceExportTaskAttributeResponse) GetBody() *UpdateResourceExportTaskAttributeResponseBody {
	return s.Body
}

func (s *UpdateResourceExportTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateResourceExportTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponse) SetStatusCode(v int32) *UpdateResourceExportTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponse) SetBody(v *UpdateResourceExportTaskAttributeResponseBody) *UpdateResourceExportTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
