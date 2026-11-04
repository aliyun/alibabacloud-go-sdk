// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceDeviceGroupsByDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceDeviceGroupsByDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceDeviceGroupsByDeviceResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceDeviceGroupsByDeviceResponseBody) *QueryFaceDeviceGroupsByDeviceResponse
	GetBody() *QueryFaceDeviceGroupsByDeviceResponseBody
}

type QueryFaceDeviceGroupsByDeviceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceDeviceGroupsByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) GetBody() *QueryFaceDeviceGroupsByDeviceResponseBody {
	return s.Body
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetHeaders(v map[string]*string) *QueryFaceDeviceGroupsByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetStatusCode(v int32) *QueryFaceDeviceGroupsByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetBody(v *QueryFaceDeviceGroupsByDeviceResponseBody) *QueryFaceDeviceGroupsByDeviceResponse {
	s.Body = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
