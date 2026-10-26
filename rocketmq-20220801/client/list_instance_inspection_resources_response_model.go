// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceInspectionResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceInspectionResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceInspectionResourcesResponseBody) *ListInstanceInspectionResourcesResponse
	GetBody() *ListInstanceInspectionResourcesResponseBody
}

type ListInstanceInspectionResourcesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceInspectionResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceInspectionResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceInspectionResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceInspectionResourcesResponse) GetBody() *ListInstanceInspectionResourcesResponseBody {
	return s.Body
}

func (s *ListInstanceInspectionResourcesResponse) SetHeaders(v map[string]*string) *ListInstanceInspectionResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceInspectionResourcesResponse) SetStatusCode(v int32) *ListInstanceInspectionResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponse) SetBody(v *ListInstanceInspectionResourcesResponseBody) *ListInstanceInspectionResourcesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceInspectionResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
