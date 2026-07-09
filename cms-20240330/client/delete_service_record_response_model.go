// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceRecordResponseBody) *DeleteServiceRecordResponse
	GetBody() *DeleteServiceRecordResponseBody
}

type DeleteServiceRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceRecordResponse) GetBody() *DeleteServiceRecordResponseBody {
	return s.Body
}

func (s *DeleteServiceRecordResponse) SetHeaders(v map[string]*string) *DeleteServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceRecordResponse) SetStatusCode(v int32) *DeleteServiceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceRecordResponse) SetBody(v *DeleteServiceRecordResponseBody) *DeleteServiceRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
