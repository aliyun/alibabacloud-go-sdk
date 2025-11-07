// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePackageStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancePackageStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancePackageStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancePackageStatesResponseBody) *ListInstancePackageStatesResponse
	GetBody() *ListInstancePackageStatesResponseBody
}

type ListInstancePackageStatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePackageStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePackageStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePackageStatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancePackageStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancePackageStatesResponse) GetBody() *ListInstancePackageStatesResponseBody {
	return s.Body
}

func (s *ListInstancePackageStatesResponse) SetHeaders(v map[string]*string) *ListInstancePackageStatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePackageStatesResponse) SetStatusCode(v int32) *ListInstancePackageStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePackageStatesResponse) SetBody(v *ListInstancePackageStatesResponseBody) *ListInstancePackageStatesResponse {
	s.Body = v
	return s
}

func (s *ListInstancePackageStatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
