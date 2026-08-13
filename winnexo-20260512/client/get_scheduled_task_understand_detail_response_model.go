// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskUnderstandDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledTaskUnderstandDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledTaskUnderstandDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledTaskUnderstandDetailResponseBody) *GetScheduledTaskUnderstandDetailResponse
	GetBody() *GetScheduledTaskUnderstandDetailResponseBody
}

type GetScheduledTaskUnderstandDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledTaskUnderstandDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledTaskUnderstandDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledTaskUnderstandDetailResponse) GetBody() *GetScheduledTaskUnderstandDetailResponseBody {
	return s.Body
}

func (s *GetScheduledTaskUnderstandDetailResponse) SetHeaders(v map[string]*string) *GetScheduledTaskUnderstandDetailResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponse) SetStatusCode(v int32) *GetScheduledTaskUnderstandDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponse) SetBody(v *GetScheduledTaskUnderstandDetailResponseBody) *GetScheduledTaskUnderstandDetailResponse {
	s.Body = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
