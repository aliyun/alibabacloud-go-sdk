// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadBookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadBookResponse
	GetStatusCode() *int32
	SetBody(v *UploadBookResponseBody) *UploadBookResponse
	GetBody() *UploadBookResponseBody
}

type UploadBookResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadBookResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadBookResponse) GoString() string {
	return s.String()
}

func (s *UploadBookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadBookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadBookResponse) GetBody() *UploadBookResponseBody {
	return s.Body
}

func (s *UploadBookResponse) SetHeaders(v map[string]*string) *UploadBookResponse {
	s.Headers = v
	return s
}

func (s *UploadBookResponse) SetStatusCode(v int32) *UploadBookResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadBookResponse) SetBody(v *UploadBookResponseBody) *UploadBookResponse {
	s.Body = v
	return s
}

func (s *UploadBookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
