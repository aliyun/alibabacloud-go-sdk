// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProvisionConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProvisionConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListProvisionConfigsOutput) *ListProvisionConfigsResponse
	GetBody() *ListProvisionConfigsOutput
}

type ListProvisionConfigsResponse struct {
	Headers    map[string]*string          `json:"headers" xml:"headers"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionConfigsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProvisionConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProvisionConfigsResponse) GetBody() *ListProvisionConfigsOutput {
	return s.Body
}

func (s *ListProvisionConfigsResponse) SetHeaders(v map[string]*string) *ListProvisionConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionConfigsResponse) SetStatusCode(v int32) *ListProvisionConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionConfigsResponse) SetBody(v *ListProvisionConfigsOutput) *ListProvisionConfigsResponse {
	s.Body = v
	return s
}

func (s *ListProvisionConfigsResponse) Validate() error {
	return dara.Validate(s)
}
