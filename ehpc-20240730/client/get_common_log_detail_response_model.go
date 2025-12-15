// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonLogDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommonLogDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommonLogDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCommonLogDetailResponseBody) *GetCommonLogDetailResponse
	GetBody() *GetCommonLogDetailResponseBody
}

type GetCommonLogDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommonLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommonLogDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommonLogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommonLogDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommonLogDetailResponse) GetBody() *GetCommonLogDetailResponseBody {
	return s.Body
}

func (s *GetCommonLogDetailResponse) SetHeaders(v map[string]*string) *GetCommonLogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCommonLogDetailResponse) SetStatusCode(v int32) *GetCommonLogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommonLogDetailResponse) SetBody(v *GetCommonLogDetailResponseBody) *GetCommonLogDetailResponse {
	s.Body = v
	return s
}

func (s *GetCommonLogDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
