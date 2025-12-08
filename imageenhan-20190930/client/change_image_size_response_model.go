// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeImageSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeImageSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeImageSizeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeImageSizeResponseBody) *ChangeImageSizeResponse
	GetBody() *ChangeImageSizeResponseBody
}

type ChangeImageSizeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeImageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeImageSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeImageSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeImageSizeResponse) GetBody() *ChangeImageSizeResponseBody {
	return s.Body
}

func (s *ChangeImageSizeResponse) SetHeaders(v map[string]*string) *ChangeImageSizeResponse {
	s.Headers = v
	return s
}

func (s *ChangeImageSizeResponse) SetStatusCode(v int32) *ChangeImageSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeImageSizeResponse) SetBody(v *ChangeImageSizeResponseBody) *ChangeImageSizeResponse {
	s.Body = v
	return s
}

func (s *ChangeImageSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
