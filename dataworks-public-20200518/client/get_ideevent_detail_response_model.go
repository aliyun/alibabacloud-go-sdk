// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIDEEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIDEEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIDEEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetIDEEventDetailResponseBody) *GetIDEEventDetailResponse
	GetBody() *GetIDEEventDetailResponseBody
}

type GetIDEEventDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIDEEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIDEEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponse) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIDEEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIDEEventDetailResponse) GetBody() *GetIDEEventDetailResponseBody {
	return s.Body
}

func (s *GetIDEEventDetailResponse) SetHeaders(v map[string]*string) *GetIDEEventDetailResponse {
	s.Headers = v
	return s
}

func (s *GetIDEEventDetailResponse) SetStatusCode(v int32) *GetIDEEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIDEEventDetailResponse) SetBody(v *GetIDEEventDetailResponseBody) *GetIDEEventDetailResponse {
	s.Body = v
	return s
}

func (s *GetIDEEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
