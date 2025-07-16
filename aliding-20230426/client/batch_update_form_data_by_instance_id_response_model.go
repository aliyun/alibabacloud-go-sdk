// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateFormDataByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateFormDataByInstanceIdResponseBody) *BatchUpdateFormDataByInstanceIdResponse
	GetBody() *BatchUpdateFormDataByInstanceIdResponseBody
}

type BatchUpdateFormDataByInstanceIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateFormDataByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateFormDataByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateFormDataByInstanceIdResponse) GetBody() *BatchUpdateFormDataByInstanceIdResponseBody {
	return s.Body
}

func (s *BatchUpdateFormDataByInstanceIdResponse) SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponse) SetStatusCode(v int32) *BatchUpdateFormDataByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponse) SetBody(v *BatchUpdateFormDataByInstanceIdResponseBody) *BatchUpdateFormDataByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponse) Validate() error {
	return dara.Validate(s)
}
