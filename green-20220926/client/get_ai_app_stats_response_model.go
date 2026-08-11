// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppStatsResponseBody) *GetAiAppStatsResponse
	GetBody() *GetAiAppStatsResponseBody
}

type GetAiAppStatsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppStatsResponse) GetBody() *GetAiAppStatsResponseBody {
	return s.Body
}

func (s *GetAiAppStatsResponse) SetHeaders(v map[string]*string) *GetAiAppStatsResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppStatsResponse) SetStatusCode(v int32) *GetAiAppStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppStatsResponse) SetBody(v *GetAiAppStatsResponseBody) *GetAiAppStatsResponse {
	s.Body = v
	return s
}

func (s *GetAiAppStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
