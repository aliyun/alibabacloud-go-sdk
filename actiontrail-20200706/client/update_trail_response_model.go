// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrailResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrailResponseBody) *UpdateTrailResponse
	GetBody() *UpdateTrailResponseBody
}

type UpdateTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrailResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrailResponse) GetBody() *UpdateTrailResponseBody {
	return s.Body
}

func (s *UpdateTrailResponse) SetHeaders(v map[string]*string) *UpdateTrailResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrailResponse) SetStatusCode(v int32) *UpdateTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrailResponse) SetBody(v *UpdateTrailResponseBody) *UpdateTrailResponse {
	s.Body = v
	return s
}

func (s *UpdateTrailResponse) Validate() error {
	return dara.Validate(s)
}
