// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupDetailResponseBody) *GetGroupDetailResponse
	GetBody() *GetGroupDetailResponseBody
}

type GetGroupDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupDetailResponse) GetBody() *GetGroupDetailResponseBody {
	return s.Body
}

func (s *GetGroupDetailResponse) SetHeaders(v map[string]*string) *GetGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGroupDetailResponse) SetStatusCode(v int32) *GetGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupDetailResponse) SetBody(v *GetGroupDetailResponseBody) *GetGroupDetailResponse {
	s.Body = v
	return s
}

func (s *GetGroupDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
