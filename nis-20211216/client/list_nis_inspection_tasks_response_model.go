// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNisInspectionTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNisInspectionTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListNisInspectionTasksResponseBody) *ListNisInspectionTasksResponse
	GetBody() *ListNisInspectionTasksResponseBody
}

type ListNisInspectionTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNisInspectionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNisInspectionTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTasksResponse) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNisInspectionTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNisInspectionTasksResponse) GetBody() *ListNisInspectionTasksResponseBody {
	return s.Body
}

func (s *ListNisInspectionTasksResponse) SetHeaders(v map[string]*string) *ListNisInspectionTasksResponse {
	s.Headers = v
	return s
}

func (s *ListNisInspectionTasksResponse) SetStatusCode(v int32) *ListNisInspectionTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNisInspectionTasksResponse) SetBody(v *ListNisInspectionTasksResponseBody) *ListNisInspectionTasksResponse {
	s.Body = v
	return s
}

func (s *ListNisInspectionTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
