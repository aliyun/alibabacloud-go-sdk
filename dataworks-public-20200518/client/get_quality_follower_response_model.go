// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityFollowerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityFollowerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityFollowerResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityFollowerResponseBody) *GetQualityFollowerResponse
	GetBody() *GetQualityFollowerResponseBody
}

type GetQualityFollowerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityFollowerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityFollowerResponse) GoString() string {
	return s.String()
}

func (s *GetQualityFollowerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityFollowerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityFollowerResponse) GetBody() *GetQualityFollowerResponseBody {
	return s.Body
}

func (s *GetQualityFollowerResponse) SetHeaders(v map[string]*string) *GetQualityFollowerResponse {
	s.Headers = v
	return s
}

func (s *GetQualityFollowerResponse) SetStatusCode(v int32) *GetQualityFollowerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityFollowerResponse) SetBody(v *GetQualityFollowerResponseBody) *GetQualityFollowerResponse {
	s.Body = v
	return s
}

func (s *GetQualityFollowerResponse) Validate() error {
	return dara.Validate(s)
}
