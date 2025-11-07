// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTaskReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNisInspectionTaskReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNisInspectionTaskReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListNisInspectionTaskReportsResponseBody) *ListNisInspectionTaskReportsResponse
	GetBody() *ListNisInspectionTaskReportsResponseBody
}

type ListNisInspectionTaskReportsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNisInspectionTaskReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNisInspectionTaskReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTaskReportsResponse) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTaskReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNisInspectionTaskReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNisInspectionTaskReportsResponse) GetBody() *ListNisInspectionTaskReportsResponseBody {
	return s.Body
}

func (s *ListNisInspectionTaskReportsResponse) SetHeaders(v map[string]*string) *ListNisInspectionTaskReportsResponse {
	s.Headers = v
	return s
}

func (s *ListNisInspectionTaskReportsResponse) SetStatusCode(v int32) *ListNisInspectionTaskReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponse) SetBody(v *ListNisInspectionTaskReportsResponseBody) *ListNisInspectionTaskReportsResponse {
	s.Body = v
	return s
}

func (s *ListNisInspectionTaskReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
