// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVerifyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceVerifyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceVerifyResultResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceVerifyResultResponseBody) *DeleteFaceVerifyResultResponse
	GetBody() *DeleteFaceVerifyResultResponseBody
}

type DeleteFaceVerifyResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceVerifyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceVerifyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceVerifyResultResponse) GetBody() *DeleteFaceVerifyResultResponseBody {
	return s.Body
}

func (s *DeleteFaceVerifyResultResponse) SetHeaders(v map[string]*string) *DeleteFaceVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceVerifyResultResponse) SetStatusCode(v int32) *DeleteFaceVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceVerifyResultResponse) SetBody(v *DeleteFaceVerifyResultResponseBody) *DeleteFaceVerifyResultResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceVerifyResultResponse) Validate() error {
	return dara.Validate(s)
}
