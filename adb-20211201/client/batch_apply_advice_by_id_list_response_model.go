// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchApplyAdviceByIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchApplyAdviceByIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchApplyAdviceByIdListResponse
	GetStatusCode() *int32
	SetBody(v *BatchApplyAdviceByIdListResponseBody) *BatchApplyAdviceByIdListResponse
	GetBody() *BatchApplyAdviceByIdListResponseBody
}

type BatchApplyAdviceByIdListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchApplyAdviceByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchApplyAdviceByIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchApplyAdviceByIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchApplyAdviceByIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchApplyAdviceByIdListResponse) GetBody() *BatchApplyAdviceByIdListResponseBody {
	return s.Body
}

func (s *BatchApplyAdviceByIdListResponse) SetHeaders(v map[string]*string) *BatchApplyAdviceByIdListResponse {
	s.Headers = v
	return s
}

func (s *BatchApplyAdviceByIdListResponse) SetStatusCode(v int32) *BatchApplyAdviceByIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchApplyAdviceByIdListResponse) SetBody(v *BatchApplyAdviceByIdListResponseBody) *BatchApplyAdviceByIdListResponse {
	s.Body = v
	return s
}

func (s *BatchApplyAdviceByIdListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
