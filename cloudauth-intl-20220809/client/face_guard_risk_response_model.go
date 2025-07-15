// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceGuardRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceGuardRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceGuardRiskResponse
	GetStatusCode() *int32
	SetBody(v *FaceGuardRiskResponseBody) *FaceGuardRiskResponse
	GetBody() *FaceGuardRiskResponseBody
}

type FaceGuardRiskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceGuardRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceGuardRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceGuardRiskResponse) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceGuardRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceGuardRiskResponse) GetBody() *FaceGuardRiskResponseBody {
	return s.Body
}

func (s *FaceGuardRiskResponse) SetHeaders(v map[string]*string) *FaceGuardRiskResponse {
	s.Headers = v
	return s
}

func (s *FaceGuardRiskResponse) SetStatusCode(v int32) *FaceGuardRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceGuardRiskResponse) SetBody(v *FaceGuardRiskResponseBody) *FaceGuardRiskResponse {
	s.Body = v
	return s
}

func (s *FaceGuardRiskResponse) Validate() error {
	return dara.Validate(s)
}
