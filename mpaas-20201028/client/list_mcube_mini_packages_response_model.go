// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeMiniPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeMiniPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeMiniPackagesResponseBody) *ListMcubeMiniPackagesResponse
	GetBody() *ListMcubeMiniPackagesResponseBody
}

type ListMcubeMiniPackagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeMiniPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeMiniPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeMiniPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeMiniPackagesResponse) GetBody() *ListMcubeMiniPackagesResponseBody {
	return s.Body
}

func (s *ListMcubeMiniPackagesResponse) SetHeaders(v map[string]*string) *ListMcubeMiniPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeMiniPackagesResponse) SetStatusCode(v int32) *ListMcubeMiniPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeMiniPackagesResponse) SetBody(v *ListMcubeMiniPackagesResponseBody) *ListMcubeMiniPackagesResponse {
	s.Body = v
	return s
}

func (s *ListMcubeMiniPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
