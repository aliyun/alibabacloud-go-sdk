// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowCooperatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowCooperatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowCooperatorsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowCooperatorsResponseBody) *ListTaskFlowCooperatorsResponse
	GetBody() *ListTaskFlowCooperatorsResponseBody
}

type ListTaskFlowCooperatorsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowCooperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowCooperatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowCooperatorsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowCooperatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowCooperatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowCooperatorsResponse) GetBody() *ListTaskFlowCooperatorsResponseBody {
	return s.Body
}

func (s *ListTaskFlowCooperatorsResponse) SetHeaders(v map[string]*string) *ListTaskFlowCooperatorsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowCooperatorsResponse) SetStatusCode(v int32) *ListTaskFlowCooperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponse) SetBody(v *ListTaskFlowCooperatorsResponseBody) *ListTaskFlowCooperatorsResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowCooperatorsResponse) Validate() error {
	return dara.Validate(s)
}
