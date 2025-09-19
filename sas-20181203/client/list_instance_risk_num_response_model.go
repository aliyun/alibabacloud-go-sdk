// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceRiskNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceRiskNumResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceRiskNumResponseBody) *ListInstanceRiskNumResponse
	GetBody() *ListInstanceRiskNumResponseBody
}

type ListInstanceRiskNumResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRiskNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRiskNumResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceRiskNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceRiskNumResponse) GetBody() *ListInstanceRiskNumResponseBody {
	return s.Body
}

func (s *ListInstanceRiskNumResponse) SetHeaders(v map[string]*string) *ListInstanceRiskNumResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRiskNumResponse) SetStatusCode(v int32) *ListInstanceRiskNumResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRiskNumResponse) SetBody(v *ListInstanceRiskNumResponseBody) *ListInstanceRiskNumResponse {
	s.Body = v
	return s
}

func (s *ListInstanceRiskNumResponse) Validate() error {
	return dara.Validate(s)
}
