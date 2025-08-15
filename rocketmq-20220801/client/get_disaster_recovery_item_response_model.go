// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *GetDisasterRecoveryItemResponseBody) *GetDisasterRecoveryItemResponse
	GetBody() *GetDisasterRecoveryItemResponseBody
}

type GetDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDisasterRecoveryItemResponse) GetBody() *GetDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *GetDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *GetDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *GetDisasterRecoveryItemResponse) SetStatusCode(v int32) *GetDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDisasterRecoveryItemResponse) SetBody(v *GetDisasterRecoveryItemResponseBody) *GetDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *GetDisasterRecoveryItemResponse) Validate() error {
	return dara.Validate(s)
}
