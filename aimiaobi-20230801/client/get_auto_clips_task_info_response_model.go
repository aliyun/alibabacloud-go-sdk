// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoClipsTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoClipsTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoClipsTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoClipsTaskInfoResponseBody) *GetAutoClipsTaskInfoResponse
	GetBody() *GetAutoClipsTaskInfoResponseBody
}

type GetAutoClipsTaskInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoClipsTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoClipsTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoClipsTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoClipsTaskInfoResponse) GetBody() *GetAutoClipsTaskInfoResponseBody {
	return s.Body
}

func (s *GetAutoClipsTaskInfoResponse) SetHeaders(v map[string]*string) *GetAutoClipsTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAutoClipsTaskInfoResponse) SetStatusCode(v int32) *GetAutoClipsTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponse) SetBody(v *GetAutoClipsTaskInfoResponseBody) *GetAutoClipsTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetAutoClipsTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
