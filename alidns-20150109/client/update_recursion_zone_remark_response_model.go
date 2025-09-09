// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionZoneRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionZoneRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionZoneRemarkResponseBody) *UpdateRecursionZoneRemarkResponse
	GetBody() *UpdateRecursionZoneRemarkResponseBody
}

type UpdateRecursionZoneRemarkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionZoneRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionZoneRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionZoneRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionZoneRemarkResponse) GetBody() *UpdateRecursionZoneRemarkResponseBody {
	return s.Body
}

func (s *UpdateRecursionZoneRemarkResponse) SetHeaders(v map[string]*string) *UpdateRecursionZoneRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionZoneRemarkResponse) SetStatusCode(v int32) *UpdateRecursionZoneRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionZoneRemarkResponse) SetBody(v *UpdateRecursionZoneRemarkResponseBody) *UpdateRecursionZoneRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionZoneRemarkResponse) Validate() error {
	return dara.Validate(s)
}
