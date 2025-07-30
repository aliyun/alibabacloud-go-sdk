// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSyncResultResponseBody) *GetSyncResultResponse
	GetBody() *GetSyncResultResponseBody
}

type GetSyncResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSyncResultResponse) GetBody() *GetSyncResultResponseBody {
	return s.Body
}

func (s *GetSyncResultResponse) SetHeaders(v map[string]*string) *GetSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetSyncResultResponse) SetStatusCode(v int32) *GetSyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyncResultResponse) SetBody(v *GetSyncResultResponseBody) *GetSyncResultResponse {
	s.Body = v
	return s
}

func (s *GetSyncResultResponse) Validate() error {
	return dara.Validate(s)
}
