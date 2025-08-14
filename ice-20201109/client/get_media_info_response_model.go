// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaInfoResponseBody) *GetMediaInfoResponse
	GetBody() *GetMediaInfoResponseBody
}

type GetMediaInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaInfoResponse) GetBody() *GetMediaInfoResponseBody {
	return s.Body
}

func (s *GetMediaInfoResponse) SetHeaders(v map[string]*string) *GetMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMediaInfoResponse) SetStatusCode(v int32) *GetMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaInfoResponse) SetBody(v *GetMediaInfoResponseBody) *GetMediaInfoResponse {
	s.Body = v
	return s
}

func (s *GetMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
