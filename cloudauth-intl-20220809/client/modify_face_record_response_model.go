// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFaceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFaceRecordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFaceRecordResponseBody) *ModifyFaceRecordResponse
	GetBody() *ModifyFaceRecordResponseBody
}

type ModifyFaceRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFaceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFaceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceRecordResponse) GoString() string {
	return s.String()
}

func (s *ModifyFaceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFaceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFaceRecordResponse) GetBody() *ModifyFaceRecordResponseBody {
	return s.Body
}

func (s *ModifyFaceRecordResponse) SetHeaders(v map[string]*string) *ModifyFaceRecordResponse {
	s.Headers = v
	return s
}

func (s *ModifyFaceRecordResponse) SetStatusCode(v int32) *ModifyFaceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFaceRecordResponse) SetBody(v *ModifyFaceRecordResponseBody) *ModifyFaceRecordResponse {
	s.Body = v
	return s
}

func (s *ModifyFaceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
