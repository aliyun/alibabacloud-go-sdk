// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceInspectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceInspectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceInspectionsResponseBody) *ListInstanceInspectionsResponse
	GetBody() *ListInstanceInspectionsResponseBody
}

type ListInstanceInspectionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceInspectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceInspectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceInspectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceInspectionsResponse) GetBody() *ListInstanceInspectionsResponseBody {
	return s.Body
}

func (s *ListInstanceInspectionsResponse) SetHeaders(v map[string]*string) *ListInstanceInspectionsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceInspectionsResponse) SetStatusCode(v int32) *ListInstanceInspectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceInspectionsResponse) SetBody(v *ListInstanceInspectionsResponseBody) *ListInstanceInspectionsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceInspectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
