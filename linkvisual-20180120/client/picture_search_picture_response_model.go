// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPictureSearchPictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PictureSearchPictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PictureSearchPictureResponse
	GetStatusCode() *int32
	SetBody(v *PictureSearchPictureResponseBody) *PictureSearchPictureResponse
	GetBody() *PictureSearchPictureResponseBody
}

type PictureSearchPictureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PictureSearchPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PictureSearchPictureResponse) String() string {
	return dara.Prettify(s)
}

func (s PictureSearchPictureResponse) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PictureSearchPictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PictureSearchPictureResponse) GetBody() *PictureSearchPictureResponseBody {
	return s.Body
}

func (s *PictureSearchPictureResponse) SetHeaders(v map[string]*string) *PictureSearchPictureResponse {
	s.Headers = v
	return s
}

func (s *PictureSearchPictureResponse) SetStatusCode(v int32) *PictureSearchPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *PictureSearchPictureResponse) SetBody(v *PictureSearchPictureResponseBody) *PictureSearchPictureResponse {
	s.Body = v
	return s
}

func (s *PictureSearchPictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
