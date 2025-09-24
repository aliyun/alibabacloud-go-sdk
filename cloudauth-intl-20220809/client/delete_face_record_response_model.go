// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceRecordResponseBody) *DeleteFaceRecordResponse
	GetBody() *DeleteFaceRecordResponseBody
}

type DeleteFaceRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceRecordResponse) GetBody() *DeleteFaceRecordResponseBody {
	return s.Body
}

func (s *DeleteFaceRecordResponse) SetHeaders(v map[string]*string) *DeleteFaceRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceRecordResponse) SetStatusCode(v int32) *DeleteFaceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceRecordResponse) SetBody(v *DeleteFaceRecordResponseBody) *DeleteFaceRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceRecordResponse) Validate() error {
	return dara.Validate(s)
}
