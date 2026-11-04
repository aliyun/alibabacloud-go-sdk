// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureLifeCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDevicePictureLifeCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDevicePictureLifeCycleResponse
	GetStatusCode() *int32
	SetBody(v *QueryDevicePictureLifeCycleResponseBody) *QueryDevicePictureLifeCycleResponse
	GetBody() *QueryDevicePictureLifeCycleResponseBody
}

type QueryDevicePictureLifeCycleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureLifeCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDevicePictureLifeCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDevicePictureLifeCycleResponse) GetBody() *QueryDevicePictureLifeCycleResponseBody {
	return s.Body
}

func (s *QueryDevicePictureLifeCycleResponse) SetHeaders(v map[string]*string) *QueryDevicePictureLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureLifeCycleResponse) SetStatusCode(v int32) *QueryDevicePictureLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponse) SetBody(v *QueryDevicePictureLifeCycleResponseBody) *QueryDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

func (s *QueryDevicePictureLifeCycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
