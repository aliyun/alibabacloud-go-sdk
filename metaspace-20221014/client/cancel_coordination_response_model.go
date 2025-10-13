// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCoordinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCoordinationResponse
	GetStatusCode() *int32
	SetBody(v *CancelCoordinationResponseBody) *CancelCoordinationResponse
	GetBody() *CancelCoordinationResponseBody
}

type CancelCoordinationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCoordinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCoordinationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationResponse) GoString() string {
	return s.String()
}

func (s *CancelCoordinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCoordinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCoordinationResponse) GetBody() *CancelCoordinationResponseBody {
	return s.Body
}

func (s *CancelCoordinationResponse) SetHeaders(v map[string]*string) *CancelCoordinationResponse {
	s.Headers = v
	return s
}

func (s *CancelCoordinationResponse) SetStatusCode(v int32) *CancelCoordinationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCoordinationResponse) SetBody(v *CancelCoordinationResponseBody) *CancelCoordinationResponse {
	s.Body = v
	return s
}

func (s *CancelCoordinationResponse) Validate() error {
	return dara.Validate(s)
}
