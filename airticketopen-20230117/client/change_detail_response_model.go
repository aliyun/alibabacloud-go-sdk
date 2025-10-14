// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDetailResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDetailResponseBody) *ChangeDetailResponse
	GetBody() *ChangeDetailResponseBody
}

type ChangeDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDetailResponse) GetBody() *ChangeDetailResponseBody {
	return s.Body
}

func (s *ChangeDetailResponse) SetHeaders(v map[string]*string) *ChangeDetailResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailResponse) SetStatusCode(v int32) *ChangeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailResponse) SetBody(v *ChangeDetailResponseBody) *ChangeDetailResponse {
	s.Body = v
	return s
}

func (s *ChangeDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
