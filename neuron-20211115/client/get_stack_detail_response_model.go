// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetStackDetailResponseBody) *GetStackDetailResponse
	GetBody() *GetStackDetailResponseBody
}

type GetStackDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStackDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackDetailResponse) GetBody() *GetStackDetailResponseBody {
	return s.Body
}

func (s *GetStackDetailResponse) SetHeaders(v map[string]*string) *GetStackDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStackDetailResponse) SetStatusCode(v int32) *GetStackDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackDetailResponse) SetBody(v *GetStackDetailResponseBody) *GetStackDetailResponse {
	s.Body = v
	return s
}

func (s *GetStackDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
