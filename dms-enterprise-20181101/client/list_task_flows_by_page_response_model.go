// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowsByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowsByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowsByPageResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowsByPageResponseBody) *ListTaskFlowsByPageResponse
	GetBody() *ListTaskFlowsByPageResponseBody
}

type ListTaskFlowsByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowsByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowsByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowsByPageResponse) GetBody() *ListTaskFlowsByPageResponseBody {
	return s.Body
}

func (s *ListTaskFlowsByPageResponse) SetHeaders(v map[string]*string) *ListTaskFlowsByPageResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowsByPageResponse) SetStatusCode(v int32) *ListTaskFlowsByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowsByPageResponse) SetBody(v *ListTaskFlowsByPageResponseBody) *ListTaskFlowsByPageResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowsByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
