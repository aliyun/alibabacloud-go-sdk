// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBizEntityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBizEntityInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBizEntityInfoResponseBody) *GetBizEntityInfoResponse
	GetBody() *GetBizEntityInfoResponseBody
}

type GetBizEntityInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBizEntityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBizEntityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBizEntityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBizEntityInfoResponse) GetBody() *GetBizEntityInfoResponseBody {
	return s.Body
}

func (s *GetBizEntityInfoResponse) SetHeaders(v map[string]*string) *GetBizEntityInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBizEntityInfoResponse) SetStatusCode(v int32) *GetBizEntityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBizEntityInfoResponse) SetBody(v *GetBizEntityInfoResponseBody) *GetBizEntityInfoResponse {
	s.Body = v
	return s
}

func (s *GetBizEntityInfoResponse) Validate() error {
	return dara.Validate(s)
}
