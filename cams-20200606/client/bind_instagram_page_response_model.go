// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstagramPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindInstagramPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindInstagramPageResponse
	GetStatusCode() *int32
	SetBody(v *BindInstagramPageResponseBody) *BindInstagramPageResponse
	GetBody() *BindInstagramPageResponseBody
}

type BindInstagramPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindInstagramPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindInstagramPageResponse) String() string {
	return dara.Prettify(s)
}

func (s BindInstagramPageResponse) GoString() string {
	return s.String()
}

func (s *BindInstagramPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindInstagramPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindInstagramPageResponse) GetBody() *BindInstagramPageResponseBody {
	return s.Body
}

func (s *BindInstagramPageResponse) SetHeaders(v map[string]*string) *BindInstagramPageResponse {
	s.Headers = v
	return s
}

func (s *BindInstagramPageResponse) SetStatusCode(v int32) *BindInstagramPageResponse {
	s.StatusCode = &v
	return s
}

func (s *BindInstagramPageResponse) SetBody(v *BindInstagramPageResponseBody) *BindInstagramPageResponse {
	s.Body = v
	return s
}

func (s *BindInstagramPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
