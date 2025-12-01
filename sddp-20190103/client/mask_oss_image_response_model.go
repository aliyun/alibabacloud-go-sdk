// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaskOssImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MaskOssImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MaskOssImageResponse
	GetStatusCode() *int32
	SetBody(v *MaskOssImageResponseBody) *MaskOssImageResponse
	GetBody() *MaskOssImageResponseBody
}

type MaskOssImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MaskOssImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MaskOssImageResponse) String() string {
	return dara.Prettify(s)
}

func (s MaskOssImageResponse) GoString() string {
	return s.String()
}

func (s *MaskOssImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MaskOssImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MaskOssImageResponse) GetBody() *MaskOssImageResponseBody {
	return s.Body
}

func (s *MaskOssImageResponse) SetHeaders(v map[string]*string) *MaskOssImageResponse {
	s.Headers = v
	return s
}

func (s *MaskOssImageResponse) SetStatusCode(v int32) *MaskOssImageResponse {
	s.StatusCode = &v
	return s
}

func (s *MaskOssImageResponse) SetBody(v *MaskOssImageResponseBody) *MaskOssImageResponse {
	s.Body = v
	return s
}

func (s *MaskOssImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
