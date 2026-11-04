// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDevicePictureLifeCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDevicePictureLifeCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDevicePictureLifeCycleResponse
	GetStatusCode() *int32
	SetBody(v *SetDevicePictureLifeCycleResponseBody) *SetDevicePictureLifeCycleResponse
	GetBody() *SetDevicePictureLifeCycleResponseBody
}

type SetDevicePictureLifeCycleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDevicePictureLifeCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDevicePictureLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDevicePictureLifeCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDevicePictureLifeCycleResponse) GetBody() *SetDevicePictureLifeCycleResponseBody {
	return s.Body
}

func (s *SetDevicePictureLifeCycleResponse) SetHeaders(v map[string]*string) *SetDevicePictureLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *SetDevicePictureLifeCycleResponse) SetStatusCode(v int32) *SetDevicePictureLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponse) SetBody(v *SetDevicePictureLifeCycleResponseBody) *SetDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

func (s *SetDevicePictureLifeCycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
