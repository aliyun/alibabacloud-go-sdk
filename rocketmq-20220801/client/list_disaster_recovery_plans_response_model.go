// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDisasterRecoveryPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDisasterRecoveryPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListDisasterRecoveryPlansResponseBody) *ListDisasterRecoveryPlansResponse
	GetBody() *ListDisasterRecoveryPlansResponseBody
}

type ListDisasterRecoveryPlansResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisasterRecoveryPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisasterRecoveryPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponse) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDisasterRecoveryPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDisasterRecoveryPlansResponse) GetBody() *ListDisasterRecoveryPlansResponseBody {
	return s.Body
}

func (s *ListDisasterRecoveryPlansResponse) SetHeaders(v map[string]*string) *ListDisasterRecoveryPlansResponse {
	s.Headers = v
	return s
}

func (s *ListDisasterRecoveryPlansResponse) SetStatusCode(v int32) *ListDisasterRecoveryPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponse) SetBody(v *ListDisasterRecoveryPlansResponseBody) *ListDisasterRecoveryPlansResponse {
	s.Body = v
	return s
}

func (s *ListDisasterRecoveryPlansResponse) Validate() error {
	return dara.Validate(s)
}
