// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdHocFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdHocFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdHocFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdHocFileResponseBody) *DeleteAdHocFileResponse
	GetBody() *DeleteAdHocFileResponseBody
}

type DeleteAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdHocFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdHocFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdHocFileResponse) GetBody() *DeleteAdHocFileResponseBody {
	return s.Body
}

func (s *DeleteAdHocFileResponse) SetHeaders(v map[string]*string) *DeleteAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdHocFileResponse) SetStatusCode(v int32) *DeleteAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdHocFileResponse) SetBody(v *DeleteAdHocFileResponseBody) *DeleteAdHocFileResponse {
	s.Body = v
	return s
}

func (s *DeleteAdHocFileResponse) Validate() error {
	return dara.Validate(s)
}
