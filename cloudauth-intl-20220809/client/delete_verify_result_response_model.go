// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVerifyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVerifyResultResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVerifyResultResponseBody) *DeleteVerifyResultResponse
	GetBody() *DeleteVerifyResultResponseBody
}

type DeleteVerifyResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVerifyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVerifyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVerifyResultResponse) GetBody() *DeleteVerifyResultResponseBody {
	return s.Body
}

func (s *DeleteVerifyResultResponse) SetHeaders(v map[string]*string) *DeleteVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifyResultResponse) SetStatusCode(v int32) *DeleteVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVerifyResultResponse) SetBody(v *DeleteVerifyResultResponseBody) *DeleteVerifyResultResponse {
	s.Body = v
	return s
}

func (s *DeleteVerifyResultResponse) Validate() error {
	return dara.Validate(s)
}
