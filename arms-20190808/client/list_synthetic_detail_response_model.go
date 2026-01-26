// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyntheticDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSyntheticDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSyntheticDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListSyntheticDetailResponseBody) *ListSyntheticDetailResponse
	GetBody() *ListSyntheticDetailResponseBody
}

type ListSyntheticDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSyntheticDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSyntheticDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailResponse) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSyntheticDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSyntheticDetailResponse) GetBody() *ListSyntheticDetailResponseBody {
	return s.Body
}

func (s *ListSyntheticDetailResponse) SetHeaders(v map[string]*string) *ListSyntheticDetailResponse {
	s.Headers = v
	return s
}

func (s *ListSyntheticDetailResponse) SetStatusCode(v int32) *ListSyntheticDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSyntheticDetailResponse) SetBody(v *ListSyntheticDetailResponseBody) *ListSyntheticDetailResponse {
	s.Body = v
	return s
}

func (s *ListSyntheticDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
