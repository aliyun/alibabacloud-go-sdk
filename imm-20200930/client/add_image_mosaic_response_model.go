// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageMosaicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageMosaicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageMosaicResponse
	GetStatusCode() *int32
	SetBody(v *AddImageMosaicResponseBody) *AddImageMosaicResponse
	GetBody() *AddImageMosaicResponseBody
}

type AddImageMosaicResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageMosaicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageMosaicResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicResponse) GoString() string {
	return s.String()
}

func (s *AddImageMosaicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageMosaicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageMosaicResponse) GetBody() *AddImageMosaicResponseBody {
	return s.Body
}

func (s *AddImageMosaicResponse) SetHeaders(v map[string]*string) *AddImageMosaicResponse {
	s.Headers = v
	return s
}

func (s *AddImageMosaicResponse) SetStatusCode(v int32) *AddImageMosaicResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageMosaicResponse) SetBody(v *AddImageMosaicResponseBody) *AddImageMosaicResponse {
	s.Body = v
	return s
}

func (s *AddImageMosaicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
