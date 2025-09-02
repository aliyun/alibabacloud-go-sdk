// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityFollowerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityFollowerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityFollowerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityFollowerResponseBody) *UpdateQualityFollowerResponse
	GetBody() *UpdateQualityFollowerResponseBody
}

type UpdateQualityFollowerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityFollowerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityFollowerResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityFollowerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityFollowerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityFollowerResponse) GetBody() *UpdateQualityFollowerResponseBody {
	return s.Body
}

func (s *UpdateQualityFollowerResponse) SetHeaders(v map[string]*string) *UpdateQualityFollowerResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityFollowerResponse) SetStatusCode(v int32) *UpdateQualityFollowerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityFollowerResponse) SetBody(v *UpdateQualityFollowerResponseBody) *UpdateQualityFollowerResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityFollowerResponse) Validate() error {
	return dara.Validate(s)
}
