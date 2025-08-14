// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublicMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublicMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPublicMediaInfoResponseBody) *GetPublicMediaInfoResponse
	GetBody() *GetPublicMediaInfoResponseBody
}

type GetPublicMediaInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublicMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublicMediaInfoResponse) GetBody() *GetPublicMediaInfoResponseBody {
	return s.Body
}

func (s *GetPublicMediaInfoResponse) SetHeaders(v map[string]*string) *GetPublicMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPublicMediaInfoResponse) SetStatusCode(v int32) *GetPublicMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicMediaInfoResponse) SetBody(v *GetPublicMediaInfoResponseBody) *GetPublicMediaInfoResponse {
	s.Body = v
	return s
}

func (s *GetPublicMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
