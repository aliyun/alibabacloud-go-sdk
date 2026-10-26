// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceInspectionRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceInspectionRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceInspectionRunsResponseBody) *ListInstanceInspectionRunsResponse
	GetBody() *ListInstanceInspectionRunsResponseBody
}

type ListInstanceInspectionRunsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceInspectionRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceInspectionRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceInspectionRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceInspectionRunsResponse) GetBody() *ListInstanceInspectionRunsResponseBody {
	return s.Body
}

func (s *ListInstanceInspectionRunsResponse) SetHeaders(v map[string]*string) *ListInstanceInspectionRunsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceInspectionRunsResponse) SetStatusCode(v int32) *ListInstanceInspectionRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceInspectionRunsResponse) SetBody(v *ListInstanceInspectionRunsResponseBody) *ListInstanceInspectionRunsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceInspectionRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
