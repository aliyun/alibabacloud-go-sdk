// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDisasterRecoveryItemResponseBody) *DeleteDisasterRecoveryItemResponse
	GetBody() *DeleteDisasterRecoveryItemResponseBody
}

type DeleteDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDisasterRecoveryItemResponse) GetBody() *DeleteDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *DeleteDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *DeleteDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteDisasterRecoveryItemResponse) SetStatusCode(v int32) *DeleteDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponse) SetBody(v *DeleteDisasterRecoveryItemResponseBody) *DeleteDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *DeleteDisasterRecoveryItemResponse) Validate() error {
	return dara.Validate(s)
}
