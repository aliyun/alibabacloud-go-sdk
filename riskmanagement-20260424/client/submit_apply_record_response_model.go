// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitApplyRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitApplyRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitApplyRecordResponse
	GetStatusCode() *int32
	SetBody(v *SubmitApplyRecordResponseBody) *SubmitApplyRecordResponse
	GetBody() *SubmitApplyRecordResponseBody
}

type SubmitApplyRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitApplyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitApplyRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitApplyRecordResponse) GoString() string {
	return s.String()
}

func (s *SubmitApplyRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitApplyRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitApplyRecordResponse) GetBody() *SubmitApplyRecordResponseBody {
	return s.Body
}

func (s *SubmitApplyRecordResponse) SetHeaders(v map[string]*string) *SubmitApplyRecordResponse {
	s.Headers = v
	return s
}

func (s *SubmitApplyRecordResponse) SetStatusCode(v int32) *SubmitApplyRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitApplyRecordResponse) SetBody(v *SubmitApplyRecordResponseBody) *SubmitApplyRecordResponse {
	s.Body = v
	return s
}

func (s *SubmitApplyRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
