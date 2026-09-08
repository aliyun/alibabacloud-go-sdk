// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskLevelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRiskLevelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRiskLevelsResponse
	GetStatusCode() *int32
	SetBody(v *ListRiskLevelsResponseBody) *ListRiskLevelsResponse
	GetBody() *ListRiskLevelsResponseBody
}

type ListRiskLevelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRiskLevelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *ListRiskLevelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRiskLevelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRiskLevelsResponse) GetBody() *ListRiskLevelsResponseBody {
	return s.Body
}

func (s *ListRiskLevelsResponse) SetHeaders(v map[string]*string) *ListRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *ListRiskLevelsResponse) SetStatusCode(v int32) *ListRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRiskLevelsResponse) SetBody(v *ListRiskLevelsResponseBody) *ListRiskLevelsResponse {
	s.Body = v
	return s
}

func (s *ListRiskLevelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
