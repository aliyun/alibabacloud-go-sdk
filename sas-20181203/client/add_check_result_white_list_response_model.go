// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckResultWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCheckResultWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCheckResultWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddCheckResultWhiteListResponseBody) *AddCheckResultWhiteListResponse
	GetBody() *AddCheckResultWhiteListResponseBody
}

type AddCheckResultWhiteListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCheckResultWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCheckResultWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCheckResultWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddCheckResultWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCheckResultWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCheckResultWhiteListResponse) GetBody() *AddCheckResultWhiteListResponseBody {
	return s.Body
}

func (s *AddCheckResultWhiteListResponse) SetHeaders(v map[string]*string) *AddCheckResultWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddCheckResultWhiteListResponse) SetStatusCode(v int32) *AddCheckResultWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCheckResultWhiteListResponse) SetBody(v *AddCheckResultWhiteListResponseBody) *AddCheckResultWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddCheckResultWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
