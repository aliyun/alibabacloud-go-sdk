// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsumeTimeRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceConsumeTimeRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceConsumeTimeRankResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceConsumeTimeRankResponseBody) *GetInstanceConsumeTimeRankResponse
	GetBody() *GetInstanceConsumeTimeRankResponseBody
}

type GetInstanceConsumeTimeRankResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceConsumeTimeRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceConsumeTimeRankResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsumeTimeRankResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceConsumeTimeRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceConsumeTimeRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceConsumeTimeRankResponse) GetBody() *GetInstanceConsumeTimeRankResponseBody {
	return s.Body
}

func (s *GetInstanceConsumeTimeRankResponse) SetHeaders(v map[string]*string) *GetInstanceConsumeTimeRankResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceConsumeTimeRankResponse) SetStatusCode(v int32) *GetInstanceConsumeTimeRankResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponse) SetBody(v *GetInstanceConsumeTimeRankResponseBody) *GetInstanceConsumeTimeRankResponse {
	s.Body = v
	return s
}

func (s *GetInstanceConsumeTimeRankResponse) Validate() error {
	return dara.Validate(s)
}
