// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceEventRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceEventRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceEventRecordResponseBody) *QueryDeviceEventRecordResponse
	GetBody() *QueryDeviceEventRecordResponseBody
}

type QueryDeviceEventRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceEventRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceEventRecordResponse) GetBody() *QueryDeviceEventRecordResponseBody {
	return s.Body
}

func (s *QueryDeviceEventRecordResponse) SetHeaders(v map[string]*string) *QueryDeviceEventRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventRecordResponse) SetStatusCode(v int32) *QueryDeviceEventRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventRecordResponse) SetBody(v *QueryDeviceEventRecordResponseBody) *QueryDeviceEventRecordResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceEventRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
