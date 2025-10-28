// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScalingConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScalingConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScalingConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListScalingConfigsOutput) *ListScalingConfigsResponse
	GetBody() *ListScalingConfigsOutput
}

type ListScalingConfigsResponse struct {
	Headers    map[string]*string        `json:"headers" xml:"headers"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScalingConfigsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScalingConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScalingConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListScalingConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScalingConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScalingConfigsResponse) GetBody() *ListScalingConfigsOutput {
	return s.Body
}

func (s *ListScalingConfigsResponse) SetHeaders(v map[string]*string) *ListScalingConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListScalingConfigsResponse) SetStatusCode(v int32) *ListScalingConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScalingConfigsResponse) SetBody(v *ListScalingConfigsOutput) *ListScalingConfigsResponse {
	s.Body = v
	return s
}

func (s *ListScalingConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
