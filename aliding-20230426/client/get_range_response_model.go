// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRangeResponse
	GetStatusCode() *int32
	SetBody(v *GetRangeResponseBody) *GetRangeResponse
	GetBody() *GetRangeResponseBody
}

type GetRangeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRangeResponse) GoString() string {
	return s.String()
}

func (s *GetRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRangeResponse) GetBody() *GetRangeResponseBody {
	return s.Body
}

func (s *GetRangeResponse) SetHeaders(v map[string]*string) *GetRangeResponse {
	s.Headers = v
	return s
}

func (s *GetRangeResponse) SetStatusCode(v int32) *GetRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRangeResponse) SetBody(v *GetRangeResponseBody) *GetRangeResponse {
	s.Body = v
	return s
}

func (s *GetRangeResponse) Validate() error {
	return dara.Validate(s)
}
