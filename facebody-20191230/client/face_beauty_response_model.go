// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceBeautyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceBeautyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceBeautyResponse
	GetStatusCode() *int32
	SetBody(v *FaceBeautyResponseBody) *FaceBeautyResponse
	GetBody() *FaceBeautyResponseBody
}

type FaceBeautyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceBeautyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceBeautyResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceBeautyResponse) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceBeautyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceBeautyResponse) GetBody() *FaceBeautyResponseBody {
	return s.Body
}

func (s *FaceBeautyResponse) SetHeaders(v map[string]*string) *FaceBeautyResponse {
	s.Headers = v
	return s
}

func (s *FaceBeautyResponse) SetStatusCode(v int32) *FaceBeautyResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceBeautyResponse) SetBody(v *FaceBeautyResponseBody) *FaceBeautyResponse {
	s.Body = v
	return s
}

func (s *FaceBeautyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
