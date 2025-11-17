// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSharedReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSharedReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListSharedReportsResponseBody) *ListSharedReportsResponse
	GetBody() *ListSharedReportsResponseBody
}

type ListSharedReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSharedReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSharedReportsResponse) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSharedReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSharedReportsResponse) GetBody() *ListSharedReportsResponseBody {
	return s.Body
}

func (s *ListSharedReportsResponse) SetHeaders(v map[string]*string) *ListSharedReportsResponse {
	s.Headers = v
	return s
}

func (s *ListSharedReportsResponse) SetStatusCode(v int32) *ListSharedReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedReportsResponse) SetBody(v *ListSharedReportsResponseBody) *ListSharedReportsResponse {
	s.Body = v
	return s
}

func (s *ListSharedReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
