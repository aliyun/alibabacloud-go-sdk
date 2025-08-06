// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachedMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttachedMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttachedMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAttachedMediaInfoResponseBody) *GetAttachedMediaInfoResponse
	GetBody() *GetAttachedMediaInfoResponseBody
}

type GetAttachedMediaInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttachedMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttachedMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttachedMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttachedMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttachedMediaInfoResponse) GetBody() *GetAttachedMediaInfoResponseBody {
	return s.Body
}

func (s *GetAttachedMediaInfoResponse) SetHeaders(v map[string]*string) *GetAttachedMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAttachedMediaInfoResponse) SetStatusCode(v int32) *GetAttachedMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachedMediaInfoResponse) SetBody(v *GetAttachedMediaInfoResponseBody) *GetAttachedMediaInfoResponse {
	s.Body = v
	return s
}

func (s *GetAttachedMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
