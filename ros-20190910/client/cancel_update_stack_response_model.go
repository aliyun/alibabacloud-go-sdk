// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpdateStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelUpdateStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelUpdateStackResponse
	GetStatusCode() *int32
	SetBody(v *CancelUpdateStackResponseBody) *CancelUpdateStackResponse
	GetBody() *CancelUpdateStackResponseBody
}

type CancelUpdateStackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelUpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelUpdateStackResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelUpdateStackResponse) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelUpdateStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelUpdateStackResponse) GetBody() *CancelUpdateStackResponseBody {
	return s.Body
}

func (s *CancelUpdateStackResponse) SetHeaders(v map[string]*string) *CancelUpdateStackResponse {
	s.Headers = v
	return s
}

func (s *CancelUpdateStackResponse) SetStatusCode(v int32) *CancelUpdateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUpdateStackResponse) SetBody(v *CancelUpdateStackResponseBody) *CancelUpdateStackResponse {
	s.Body = v
	return s
}

func (s *CancelUpdateStackResponse) Validate() error {
	return dara.Validate(s)
}
