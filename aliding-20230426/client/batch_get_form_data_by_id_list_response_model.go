// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetFormDataByIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetFormDataByIdListResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetFormDataByIdListResponseBody) *BatchGetFormDataByIdListResponse
	GetBody() *BatchGetFormDataByIdListResponseBody
}

type BatchGetFormDataByIdListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetFormDataByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetFormDataByIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetFormDataByIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetFormDataByIdListResponse) GetBody() *BatchGetFormDataByIdListResponseBody {
	return s.Body
}

func (s *BatchGetFormDataByIdListResponse) SetHeaders(v map[string]*string) *BatchGetFormDataByIdListResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFormDataByIdListResponse) SetStatusCode(v int32) *BatchGetFormDataByIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetFormDataByIdListResponse) SetBody(v *BatchGetFormDataByIdListResponseBody) *BatchGetFormDataByIdListResponse {
	s.Body = v
	return s
}

func (s *BatchGetFormDataByIdListResponse) Validate() error {
	return dara.Validate(s)
}
