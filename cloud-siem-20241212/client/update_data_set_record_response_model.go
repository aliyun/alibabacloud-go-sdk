// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSetRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSetRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSetRecordResponseBody) *UpdateDataSetRecordResponse
	GetBody() *UpdateDataSetRecordResponseBody
}

type UpdateDataSetRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSetRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSetRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSetRecordResponse) GetBody() *UpdateDataSetRecordResponseBody {
	return s.Body
}

func (s *UpdateDataSetRecordResponse) SetHeaders(v map[string]*string) *UpdateDataSetRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSetRecordResponse) SetStatusCode(v int32) *UpdateDataSetRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSetRecordResponse) SetBody(v *UpdateDataSetRecordResponseBody) *UpdateDataSetRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSetRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
