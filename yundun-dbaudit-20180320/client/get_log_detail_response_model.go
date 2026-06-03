// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetLogDetailResponseBody) *GetLogDetailResponse
	GetBody() *GetLogDetailResponseBody
}

type GetLogDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLogDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogDetailResponse) GetBody() *GetLogDetailResponseBody {
	return s.Body
}

func (s *GetLogDetailResponse) SetHeaders(v map[string]*string) *GetLogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLogDetailResponse) SetStatusCode(v int32) *GetLogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogDetailResponse) SetBody(v *GetLogDetailResponseBody) *GetLogDetailResponse {
	s.Body = v
	return s
}

func (s *GetLogDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
