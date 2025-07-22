// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKillInstanceSessionTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKillInstanceSessionTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKillInstanceSessionTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetKillInstanceSessionTaskResultResponseBody) *GetKillInstanceSessionTaskResultResponse
	GetBody() *GetKillInstanceSessionTaskResultResponseBody
}

type GetKillInstanceSessionTaskResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKillInstanceSessionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKillInstanceSessionTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKillInstanceSessionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetKillInstanceSessionTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKillInstanceSessionTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKillInstanceSessionTaskResultResponse) GetBody() *GetKillInstanceSessionTaskResultResponseBody {
	return s.Body
}

func (s *GetKillInstanceSessionTaskResultResponse) SetHeaders(v map[string]*string) *GetKillInstanceSessionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponse) SetStatusCode(v int32) *GetKillInstanceSessionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponse) SetBody(v *GetKillInstanceSessionTaskResultResponseBody) *GetKillInstanceSessionTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
