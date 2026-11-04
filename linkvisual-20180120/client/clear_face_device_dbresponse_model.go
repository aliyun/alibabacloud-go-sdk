// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearFaceDeviceDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearFaceDeviceDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearFaceDeviceDBResponse
	GetStatusCode() *int32
	SetBody(v *ClearFaceDeviceDBResponseBody) *ClearFaceDeviceDBResponse
	GetBody() *ClearFaceDeviceDBResponseBody
}

type ClearFaceDeviceDBResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearFaceDeviceDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearFaceDeviceDBResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearFaceDeviceDBResponse) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearFaceDeviceDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearFaceDeviceDBResponse) GetBody() *ClearFaceDeviceDBResponseBody {
	return s.Body
}

func (s *ClearFaceDeviceDBResponse) SetHeaders(v map[string]*string) *ClearFaceDeviceDBResponse {
	s.Headers = v
	return s
}

func (s *ClearFaceDeviceDBResponse) SetStatusCode(v int32) *ClearFaceDeviceDBResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearFaceDeviceDBResponse) SetBody(v *ClearFaceDeviceDBResponseBody) *ClearFaceDeviceDBResponse {
	s.Body = v
	return s
}

func (s *ClearFaceDeviceDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
