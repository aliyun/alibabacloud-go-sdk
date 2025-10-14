// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRplInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRplInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRplInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRplInspectionTaskResponseBody) *CreateRplInspectionTaskResponse
	GetBody() *CreateRplInspectionTaskResponseBody
}

type CreateRplInspectionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRplInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRplInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRplInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRplInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRplInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRplInspectionTaskResponse) GetBody() *CreateRplInspectionTaskResponseBody {
	return s.Body
}

func (s *CreateRplInspectionTaskResponse) SetHeaders(v map[string]*string) *CreateRplInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRplInspectionTaskResponse) SetStatusCode(v int32) *CreateRplInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRplInspectionTaskResponse) SetBody(v *CreateRplInspectionTaskResponseBody) *CreateRplInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRplInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
