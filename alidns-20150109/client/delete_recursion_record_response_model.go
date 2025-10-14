// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecursionRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecursionRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecursionRecordResponseBody) *DeleteRecursionRecordResponse
	GetBody() *DeleteRecursionRecordResponseBody
}

type DeleteRecursionRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecursionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecursionRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecursionRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecursionRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecursionRecordResponse) GetBody() *DeleteRecursionRecordResponseBody {
	return s.Body
}

func (s *DeleteRecursionRecordResponse) SetHeaders(v map[string]*string) *DeleteRecursionRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecursionRecordResponse) SetStatusCode(v int32) *DeleteRecursionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecursionRecordResponse) SetBody(v *DeleteRecursionRecordResponseBody) *DeleteRecursionRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteRecursionRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
