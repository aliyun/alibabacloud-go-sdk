// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateInspectionTaskResponseBody) *CreateInspectionTaskResponse
	GetBody() *CreateInspectionTaskResponseBody
}

type CreateInspectionTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInspectionTaskResponse) GetBody() *CreateInspectionTaskResponseBody {
	return s.Body
}

func (s *CreateInspectionTaskResponse) SetHeaders(v map[string]*string) *CreateInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInspectionTaskResponse) SetStatusCode(v int32) *CreateInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInspectionTaskResponse) SetBody(v *CreateInspectionTaskResponseBody) *CreateInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
