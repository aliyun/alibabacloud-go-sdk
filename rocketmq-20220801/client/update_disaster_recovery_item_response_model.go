// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDisasterRecoveryItemResponseBody) *UpdateDisasterRecoveryItemResponse
	GetBody() *UpdateDisasterRecoveryItemResponseBody
}

type UpdateDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDisasterRecoveryItemResponse) GetBody() *UpdateDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *UpdateDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *UpdateDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateDisasterRecoveryItemResponse) SetStatusCode(v int32) *UpdateDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponse) SetBody(v *UpdateDisasterRecoveryItemResponseBody) *UpdateDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *UpdateDisasterRecoveryItemResponse) Validate() error {
	return dara.Validate(s)
}
