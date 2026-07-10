// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitPreBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSubmitPreBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSubmitPreBillResponse
	GetStatusCode() *int32
	SetBody(v *BatchSubmitPreBillResponseBody) *BatchSubmitPreBillResponse
	GetBody() *BatchSubmitPreBillResponseBody
}

type BatchSubmitPreBillResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSubmitPreBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSubmitPreBillResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillResponse) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSubmitPreBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSubmitPreBillResponse) GetBody() *BatchSubmitPreBillResponseBody {
	return s.Body
}

func (s *BatchSubmitPreBillResponse) SetHeaders(v map[string]*string) *BatchSubmitPreBillResponse {
	s.Headers = v
	return s
}

func (s *BatchSubmitPreBillResponse) SetStatusCode(v int32) *BatchSubmitPreBillResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSubmitPreBillResponse) SetBody(v *BatchSubmitPreBillResponseBody) *BatchSubmitPreBillResponse {
	s.Body = v
	return s
}

func (s *BatchSubmitPreBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
