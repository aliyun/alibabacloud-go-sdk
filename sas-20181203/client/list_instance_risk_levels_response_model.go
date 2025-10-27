// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskLevelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceRiskLevelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceRiskLevelsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceRiskLevelsResponseBody) *ListInstanceRiskLevelsResponse
	GetBody() *ListInstanceRiskLevelsResponseBody
}

type ListInstanceRiskLevelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRiskLevelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskLevelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceRiskLevelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceRiskLevelsResponse) GetBody() *ListInstanceRiskLevelsResponseBody {
	return s.Body
}

func (s *ListInstanceRiskLevelsResponse) SetHeaders(v map[string]*string) *ListInstanceRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRiskLevelsResponse) SetStatusCode(v int32) *ListInstanceRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRiskLevelsResponse) SetBody(v *ListInstanceRiskLevelsResponseBody) *ListInstanceRiskLevelsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceRiskLevelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
