// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceRecordLifeCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDeviceRecordLifeCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDeviceRecordLifeCycleResponse
	GetStatusCode() *int32
	SetBody(v *SetDeviceRecordLifeCycleResponseBody) *SetDeviceRecordLifeCycleResponse
	GetBody() *SetDeviceRecordLifeCycleResponseBody
}

type SetDeviceRecordLifeCycleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceRecordLifeCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceRecordLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDeviceRecordLifeCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDeviceRecordLifeCycleResponse) GetBody() *SetDeviceRecordLifeCycleResponseBody {
	return s.Body
}

func (s *SetDeviceRecordLifeCycleResponse) SetHeaders(v map[string]*string) *SetDeviceRecordLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceRecordLifeCycleResponse) SetStatusCode(v int32) *SetDeviceRecordLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponse) SetBody(v *SetDeviceRecordLifeCycleResponseBody) *SetDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

func (s *SetDeviceRecordLifeCycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
