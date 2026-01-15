// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IncreaseListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IncreaseListResponse
	GetStatusCode() *int32
	SetBody(v *IncreaseListResponseBody) *IncreaseListResponse
	GetBody() *IncreaseListResponseBody
}

type IncreaseListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncreaseListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncreaseListResponse) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListResponse) GoString() string {
	return s.String()
}

func (s *IncreaseListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IncreaseListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IncreaseListResponse) GetBody() *IncreaseListResponseBody {
	return s.Body
}

func (s *IncreaseListResponse) SetHeaders(v map[string]*string) *IncreaseListResponse {
	s.Headers = v
	return s
}

func (s *IncreaseListResponse) SetStatusCode(v int32) *IncreaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseListResponse) SetBody(v *IncreaseListResponseBody) *IncreaseListResponse {
	s.Body = v
	return s
}

func (s *IncreaseListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
