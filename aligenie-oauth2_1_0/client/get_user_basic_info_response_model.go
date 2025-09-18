// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserBasicInfoResponseBody) *GetUserBasicInfoResponse
	GetBody() *GetUserBasicInfoResponseBody
}

type GetUserBasicInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserBasicInfoResponse) GetBody() *GetUserBasicInfoResponseBody {
	return s.Body
}

func (s *GetUserBasicInfoResponse) SetHeaders(v map[string]*string) *GetUserBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserBasicInfoResponse) SetStatusCode(v int32) *GetUserBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBasicInfoResponse) SetBody(v *GetUserBasicInfoResponseBody) *GetUserBasicInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
