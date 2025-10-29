// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShowIntoShowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddShowIntoShowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddShowIntoShowListResponse
	GetStatusCode() *int32
	SetBody(v *AddShowIntoShowListResponseBody) *AddShowIntoShowListResponse
	GetBody() *AddShowIntoShowListResponseBody
}

type AddShowIntoShowListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShowIntoShowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddShowIntoShowListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddShowIntoShowListResponse) GoString() string {
	return s.String()
}

func (s *AddShowIntoShowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddShowIntoShowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddShowIntoShowListResponse) GetBody() *AddShowIntoShowListResponseBody {
	return s.Body
}

func (s *AddShowIntoShowListResponse) SetHeaders(v map[string]*string) *AddShowIntoShowListResponse {
	s.Headers = v
	return s
}

func (s *AddShowIntoShowListResponse) SetStatusCode(v int32) *AddShowIntoShowListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShowIntoShowListResponse) SetBody(v *AddShowIntoShowListResponseBody) *AddShowIntoShowListResponse {
	s.Body = v
	return s
}

func (s *AddShowIntoShowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
