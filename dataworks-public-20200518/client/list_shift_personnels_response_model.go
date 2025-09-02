// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftPersonnelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShiftPersonnelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShiftPersonnelsResponse
	GetStatusCode() *int32
	SetBody(v *ListShiftPersonnelsResponseBody) *ListShiftPersonnelsResponse
	GetBody() *ListShiftPersonnelsResponseBody
}

type ListShiftPersonnelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShiftPersonnelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShiftPersonnelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShiftPersonnelsResponse) GoString() string {
	return s.String()
}

func (s *ListShiftPersonnelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShiftPersonnelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShiftPersonnelsResponse) GetBody() *ListShiftPersonnelsResponseBody {
	return s.Body
}

func (s *ListShiftPersonnelsResponse) SetHeaders(v map[string]*string) *ListShiftPersonnelsResponse {
	s.Headers = v
	return s
}

func (s *ListShiftPersonnelsResponse) SetStatusCode(v int32) *ListShiftPersonnelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShiftPersonnelsResponse) SetBody(v *ListShiftPersonnelsResponseBody) *ListShiftPersonnelsResponse {
	s.Body = v
	return s
}

func (s *ListShiftPersonnelsResponse) Validate() error {
	return dara.Validate(s)
}
