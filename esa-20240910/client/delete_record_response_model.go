// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecordResponseBody) *DeleteRecordResponse
	GetBody() *DeleteRecordResponseBody
}

type DeleteRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecordResponse) GetBody() *DeleteRecordResponseBody {
	return s.Body
}

func (s *DeleteRecordResponse) SetHeaders(v map[string]*string) *DeleteRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordResponse) SetStatusCode(v int32) *DeleteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordResponse) SetBody(v *DeleteRecordResponseBody) *DeleteRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteRecordResponse) Validate() error {
	return dara.Validate(s)
}
