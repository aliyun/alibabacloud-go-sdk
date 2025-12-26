// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceRecordResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceRecordResponseBody) *AddFaceRecordResponse
	GetBody() *AddFaceRecordResponseBody
}

type AddFaceRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRecordResponse) GoString() string {
	return s.String()
}

func (s *AddFaceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceRecordResponse) GetBody() *AddFaceRecordResponseBody {
	return s.Body
}

func (s *AddFaceRecordResponse) SetHeaders(v map[string]*string) *AddFaceRecordResponse {
	s.Headers = v
	return s
}

func (s *AddFaceRecordResponse) SetStatusCode(v int32) *AddFaceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceRecordResponse) SetBody(v *AddFaceRecordResponseBody) *AddFaceRecordResponse {
	s.Body = v
	return s
}

func (s *AddFaceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
