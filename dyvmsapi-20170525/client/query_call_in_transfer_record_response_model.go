// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInTransferRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallInTransferRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallInTransferRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallInTransferRecordResponseBody) *QueryCallInTransferRecordResponse
	GetBody() *QueryCallInTransferRecordResponseBody
}

type QueryCallInTransferRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallInTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallInTransferRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInTransferRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallInTransferRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallInTransferRecordResponse) GetBody() *QueryCallInTransferRecordResponseBody {
	return s.Body
}

func (s *QueryCallInTransferRecordResponse) SetHeaders(v map[string]*string) *QueryCallInTransferRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryCallInTransferRecordResponse) SetStatusCode(v int32) *QueryCallInTransferRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallInTransferRecordResponse) SetBody(v *QueryCallInTransferRecordResponseBody) *QueryCallInTransferRecordResponse {
	s.Body = v
	return s
}

func (s *QueryCallInTransferRecordResponse) Validate() error {
	return dara.Validate(s)
}
