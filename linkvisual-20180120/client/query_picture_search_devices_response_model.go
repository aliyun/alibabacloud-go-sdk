// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureSearchDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureSearchDevicesResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureSearchDevicesResponseBody) *QueryPictureSearchDevicesResponse
	GetBody() *QueryPictureSearchDevicesResponseBody
}

type QueryPictureSearchDevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchDevicesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureSearchDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureSearchDevicesResponse) GetBody() *QueryPictureSearchDevicesResponseBody {
	return s.Body
}

func (s *QueryPictureSearchDevicesResponse) SetHeaders(v map[string]*string) *QueryPictureSearchDevicesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchDevicesResponse) SetStatusCode(v int32) *QueryPictureSearchDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchDevicesResponse) SetBody(v *QueryPictureSearchDevicesResponseBody) *QueryPictureSearchDevicesResponse {
	s.Body = v
	return s
}

func (s *QueryPictureSearchDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
