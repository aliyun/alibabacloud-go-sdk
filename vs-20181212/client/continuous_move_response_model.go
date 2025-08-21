// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousMoveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinuousMoveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinuousMoveResponse
	GetStatusCode() *int32
	SetBody(v *ContinuousMoveResponseBody) *ContinuousMoveResponse
	GetBody() *ContinuousMoveResponseBody
}

type ContinuousMoveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinuousMoveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinuousMoveResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinuousMoveResponse) GoString() string {
	return s.String()
}

func (s *ContinuousMoveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinuousMoveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinuousMoveResponse) GetBody() *ContinuousMoveResponseBody {
	return s.Body
}

func (s *ContinuousMoveResponse) SetHeaders(v map[string]*string) *ContinuousMoveResponse {
	s.Headers = v
	return s
}

func (s *ContinuousMoveResponse) SetStatusCode(v int32) *ContinuousMoveResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinuousMoveResponse) SetBody(v *ContinuousMoveResponseBody) *ContinuousMoveResponse {
	s.Body = v
	return s
}

func (s *ContinuousMoveResponse) Validate() error {
	return dara.Validate(s)
}
