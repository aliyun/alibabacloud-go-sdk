// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterTypeResponseBody) *ListClusterTypeResponse
	GetBody() *ListClusterTypeResponseBody
}

type ListClusterTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypeResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterTypeResponse) GetBody() *ListClusterTypeResponseBody {
	return s.Body
}

func (s *ListClusterTypeResponse) SetHeaders(v map[string]*string) *ListClusterTypeResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTypeResponse) SetStatusCode(v int32) *ListClusterTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterTypeResponse) SetBody(v *ListClusterTypeResponseBody) *ListClusterTypeResponse {
	s.Body = v
	return s
}

func (s *ListClusterTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
