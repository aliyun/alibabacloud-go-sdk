// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNisInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNisInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNisInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNisInspectionTaskResponseBody) *UpdateNisInspectionTaskResponse
	GetBody() *UpdateNisInspectionTaskResponseBody
}

type UpdateNisInspectionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNisInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNisInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNisInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateNisInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNisInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNisInspectionTaskResponse) GetBody() *UpdateNisInspectionTaskResponseBody {
	return s.Body
}

func (s *UpdateNisInspectionTaskResponse) SetHeaders(v map[string]*string) *UpdateNisInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateNisInspectionTaskResponse) SetStatusCode(v int32) *UpdateNisInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNisInspectionTaskResponse) SetBody(v *UpdateNisInspectionTaskResponseBody) *UpdateNisInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateNisInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
