// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowConstantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowConstantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowConstantsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowConstantsResponseBody) *ListTaskFlowConstantsResponse
	GetBody() *ListTaskFlowConstantsResponseBody
}

type ListTaskFlowConstantsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowConstantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowConstantsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowConstantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowConstantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowConstantsResponse) GetBody() *ListTaskFlowConstantsResponseBody {
	return s.Body
}

func (s *ListTaskFlowConstantsResponse) SetHeaders(v map[string]*string) *ListTaskFlowConstantsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowConstantsResponse) SetStatusCode(v int32) *ListTaskFlowConstantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowConstantsResponse) SetBody(v *ListTaskFlowConstantsResponseBody) *ListTaskFlowConstantsResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowConstantsResponse) Validate() error {
	return dara.Validate(s)
}
