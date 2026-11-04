// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceRecordLifeCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceRecordLifeCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceRecordLifeCycleResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceRecordLifeCycleResponseBody) *QueryDeviceRecordLifeCycleResponse
	GetBody() *QueryDeviceRecordLifeCycleResponseBody
}

type QueryDeviceRecordLifeCycleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceRecordLifeCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceRecordLifeCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceRecordLifeCycleResponse) GetBody() *QueryDeviceRecordLifeCycleResponseBody {
	return s.Body
}

func (s *QueryDeviceRecordLifeCycleResponse) SetHeaders(v map[string]*string) *QueryDeviceRecordLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponse) SetStatusCode(v int32) *QueryDeviceRecordLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponse) SetBody(v *QueryDeviceRecordLifeCycleResponseBody) *QueryDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
