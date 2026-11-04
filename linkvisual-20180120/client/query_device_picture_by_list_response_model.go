// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureByListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDevicePictureByListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDevicePictureByListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDevicePictureByListResponseBody) *QueryDevicePictureByListResponse
	GetBody() *QueryDevicePictureByListResponseBody
}

type QueryDevicePictureByListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureByListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureByListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureByListResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDevicePictureByListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDevicePictureByListResponse) GetBody() *QueryDevicePictureByListResponseBody {
	return s.Body
}

func (s *QueryDevicePictureByListResponse) SetHeaders(v map[string]*string) *QueryDevicePictureByListResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureByListResponse) SetStatusCode(v int32) *QueryDevicePictureByListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureByListResponse) SetBody(v *QueryDevicePictureByListResponseBody) *QueryDevicePictureByListResponse {
	s.Body = v
	return s
}

func (s *QueryDevicePictureByListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
