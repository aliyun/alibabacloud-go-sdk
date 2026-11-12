// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetImageByIdResponseBody) *GetImageByIdResponse
	GetBody() *GetImageByIdResponseBody
}

type GetImageByIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdResponse) GoString() string {
	return s.String()
}

func (s *GetImageByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageByIdResponse) GetBody() *GetImageByIdResponseBody {
	return s.Body
}

func (s *GetImageByIdResponse) SetHeaders(v map[string]*string) *GetImageByIdResponse {
	s.Headers = v
	return s
}

func (s *GetImageByIdResponse) SetStatusCode(v int32) *GetImageByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageByIdResponse) SetBody(v *GetImageByIdResponseBody) *GetImageByIdResponse {
	s.Body = v
	return s
}

func (s *GetImageByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
