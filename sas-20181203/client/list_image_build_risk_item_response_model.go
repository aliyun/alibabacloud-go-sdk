// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBuildRiskItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageBuildRiskItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageBuildRiskItemResponse
	GetStatusCode() *int32
	SetBody(v *ListImageBuildRiskItemResponseBody) *ListImageBuildRiskItemResponse
	GetBody() *ListImageBuildRiskItemResponseBody
}

type ListImageBuildRiskItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageBuildRiskItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageBuildRiskItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageBuildRiskItemResponse) GoString() string {
	return s.String()
}

func (s *ListImageBuildRiskItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageBuildRiskItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageBuildRiskItemResponse) GetBody() *ListImageBuildRiskItemResponseBody {
	return s.Body
}

func (s *ListImageBuildRiskItemResponse) SetHeaders(v map[string]*string) *ListImageBuildRiskItemResponse {
	s.Headers = v
	return s
}

func (s *ListImageBuildRiskItemResponse) SetStatusCode(v int32) *ListImageBuildRiskItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageBuildRiskItemResponse) SetBody(v *ListImageBuildRiskItemResponseBody) *ListImageBuildRiskItemResponse {
	s.Body = v
	return s
}

func (s *ListImageBuildRiskItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
