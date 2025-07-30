// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityCheckSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityCheckSchemeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityCheckSchemeResponseBody) *UpdateQualityCheckSchemeResponse
	GetBody() *UpdateQualityCheckSchemeResponseBody
}

type UpdateQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityCheckSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityCheckSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityCheckSchemeResponse) GetBody() *UpdateQualityCheckSchemeResponseBody {
	return s.Body
}

func (s *UpdateQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *UpdateQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityCheckSchemeResponse) SetStatusCode(v int32) *UpdateQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponse) SetBody(v *UpdateQualityCheckSchemeResponseBody) *UpdateQualityCheckSchemeResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityCheckSchemeResponse) Validate() error {
	return dara.Validate(s)
}
