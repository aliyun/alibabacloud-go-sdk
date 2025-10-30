// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInspectReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterInspectReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterInspectReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterInspectReportsResponseBody) *ListClusterInspectReportsResponse
	GetBody() *ListClusterInspectReportsResponseBody
}

type ListClusterInspectReportsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterInspectReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterInspectReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInspectReportsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterInspectReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterInspectReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterInspectReportsResponse) GetBody() *ListClusterInspectReportsResponseBody {
	return s.Body
}

func (s *ListClusterInspectReportsResponse) SetHeaders(v map[string]*string) *ListClusterInspectReportsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterInspectReportsResponse) SetStatusCode(v int32) *ListClusterInspectReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterInspectReportsResponse) SetBody(v *ListClusterInspectReportsResponseBody) *ListClusterInspectReportsResponse {
	s.Body = v
	return s
}

func (s *ListClusterInspectReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
