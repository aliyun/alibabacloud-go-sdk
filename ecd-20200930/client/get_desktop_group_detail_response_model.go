// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesktopGroupDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDesktopGroupDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDesktopGroupDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDesktopGroupDetailResponseBody) *GetDesktopGroupDetailResponse
	GetBody() *GetDesktopGroupDetailResponseBody
}

type GetDesktopGroupDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDesktopGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDesktopGroupDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDesktopGroupDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDesktopGroupDetailResponse) GetBody() *GetDesktopGroupDetailResponseBody {
	return s.Body
}

func (s *GetDesktopGroupDetailResponse) SetHeaders(v map[string]*string) *GetDesktopGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDesktopGroupDetailResponse) SetStatusCode(v int32) *GetDesktopGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDesktopGroupDetailResponse) SetBody(v *GetDesktopGroupDetailResponseBody) *GetDesktopGroupDetailResponse {
	s.Body = v
	return s
}

func (s *GetDesktopGroupDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
