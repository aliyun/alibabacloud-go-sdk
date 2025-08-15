// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *AddDisasterRecoveryItemResponseBody) *AddDisasterRecoveryItemResponse
	GetBody() *AddDisasterRecoveryItemResponseBody
}

type AddDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDisasterRecoveryItemResponse) GetBody() *AddDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *AddDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *AddDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *AddDisasterRecoveryItemResponse) SetStatusCode(v int32) *AddDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponse) SetBody(v *AddDisasterRecoveryItemResponseBody) *AddDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *AddDisasterRecoveryItemResponse) Validate() error {
	return dara.Validate(s)
}
