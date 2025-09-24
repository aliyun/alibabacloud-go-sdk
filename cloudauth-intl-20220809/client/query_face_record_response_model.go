// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceRecordResponseBody) *QueryFaceRecordResponse
	GetBody() *QueryFaceRecordResponseBody
}

type QueryFaceRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceRecordResponse) GetBody() *QueryFaceRecordResponseBody {
	return s.Body
}

func (s *QueryFaceRecordResponse) SetHeaders(v map[string]*string) *QueryFaceRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceRecordResponse) SetStatusCode(v int32) *QueryFaceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceRecordResponse) SetBody(v *QueryFaceRecordResponseBody) *QueryFaceRecordResponse {
	s.Body = v
	return s
}

func (s *QueryFaceRecordResponse) Validate() error {
	return dara.Validate(s)
}
