// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSowListResponse
	GetStatusCode() *int32
	SetBody(v *GetSowListResponseBody) *GetSowListResponse
	GetBody() *GetSowListResponseBody
}

type GetSowListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSowListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSowListResponse) GoString() string {
	return s.String()
}

func (s *GetSowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSowListResponse) GetBody() *GetSowListResponseBody {
	return s.Body
}

func (s *GetSowListResponse) SetHeaders(v map[string]*string) *GetSowListResponse {
	s.Headers = v
	return s
}

func (s *GetSowListResponse) SetStatusCode(v int32) *GetSowListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSowListResponse) SetBody(v *GetSowListResponseBody) *GetSowListResponse {
	s.Body = v
	return s
}

func (s *GetSowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
