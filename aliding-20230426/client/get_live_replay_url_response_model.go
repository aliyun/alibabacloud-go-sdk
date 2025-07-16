// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveReplayUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveReplayUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveReplayUrlResponseBody) *GetLiveReplayUrlResponse
	GetBody() *GetLiveReplayUrlResponseBody
}

type GetLiveReplayUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveReplayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveReplayUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlResponse) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveReplayUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveReplayUrlResponse) GetBody() *GetLiveReplayUrlResponseBody {
	return s.Body
}

func (s *GetLiveReplayUrlResponse) SetHeaders(v map[string]*string) *GetLiveReplayUrlResponse {
	s.Headers = v
	return s
}

func (s *GetLiveReplayUrlResponse) SetStatusCode(v int32) *GetLiveReplayUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveReplayUrlResponse) SetBody(v *GetLiveReplayUrlResponseBody) *GetLiveReplayUrlResponse {
	s.Body = v
	return s
}

func (s *GetLiveReplayUrlResponse) Validate() error {
	return dara.Validate(s)
}
