// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRowsVisibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRowsVisibilityResponse
	GetStatusCode() *int32
	SetBody(v *SetRowsVisibilityResponseBody) *SetRowsVisibilityResponse
	GetBody() *SetRowsVisibilityResponseBody
}

type SetRowsVisibilityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRowsVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRowsVisibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityResponse) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRowsVisibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRowsVisibilityResponse) GetBody() *SetRowsVisibilityResponseBody {
	return s.Body
}

func (s *SetRowsVisibilityResponse) SetHeaders(v map[string]*string) *SetRowsVisibilityResponse {
	s.Headers = v
	return s
}

func (s *SetRowsVisibilityResponse) SetStatusCode(v int32) *SetRowsVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRowsVisibilityResponse) SetBody(v *SetRowsVisibilityResponseBody) *SetRowsVisibilityResponse {
	s.Body = v
	return s
}

func (s *SetRowsVisibilityResponse) Validate() error {
	return dara.Validate(s)
}
