// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncInstanceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDISyncInstanceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDISyncInstanceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDISyncInstanceInfoResponseBody) *GetDISyncInstanceInfoResponse
	GetBody() *GetDISyncInstanceInfoResponseBody
}

type GetDISyncInstanceInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDISyncInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDISyncInstanceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDISyncInstanceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDISyncInstanceInfoResponse) GetBody() *GetDISyncInstanceInfoResponseBody {
	return s.Body
}

func (s *GetDISyncInstanceInfoResponse) SetHeaders(v map[string]*string) *GetDISyncInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDISyncInstanceInfoResponse) SetStatusCode(v int32) *GetDISyncInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDISyncInstanceInfoResponse) SetBody(v *GetDISyncInstanceInfoResponseBody) *GetDISyncInstanceInfoResponse {
	s.Body = v
	return s
}

func (s *GetDISyncInstanceInfoResponse) Validate() error {
	return dara.Validate(s)
}
