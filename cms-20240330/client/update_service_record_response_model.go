// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceRecordResponseBody) *UpdateServiceRecordResponse
	GetBody() *UpdateServiceRecordResponseBody
}

type UpdateServiceRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceRecordResponse) GetBody() *UpdateServiceRecordResponseBody {
	return s.Body
}

func (s *UpdateServiceRecordResponse) SetHeaders(v map[string]*string) *UpdateServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceRecordResponse) SetStatusCode(v int32) *UpdateServiceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceRecordResponse) SetBody(v *UpdateServiceRecordResponseBody) *UpdateServiceRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
