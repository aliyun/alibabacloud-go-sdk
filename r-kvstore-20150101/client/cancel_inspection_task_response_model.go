// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelInspectionTaskResponseBody) *CancelInspectionTaskResponse
	GetBody() *CancelInspectionTaskResponseBody
}

type CancelInspectionTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelInspectionTaskResponse) GetBody() *CancelInspectionTaskResponseBody {
	return s.Body
}

func (s *CancelInspectionTaskResponse) SetHeaders(v map[string]*string) *CancelInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelInspectionTaskResponse) SetStatusCode(v int32) *CancelInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelInspectionTaskResponse) SetBody(v *CancelInspectionTaskResponseBody) *CancelInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *CancelInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
