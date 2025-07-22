// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestOriginStatByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFullRequestOriginStatByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFullRequestOriginStatByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetFullRequestOriginStatByInstanceIdResponseBody) *GetFullRequestOriginStatByInstanceIdResponse
	GetBody() *GetFullRequestOriginStatByInstanceIdResponseBody
}

type GetFullRequestOriginStatByInstanceIdResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFullRequestOriginStatByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) GetBody() *GetFullRequestOriginStatByInstanceIdResponseBody {
	return s.Body
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) SetHeaders(v map[string]*string) *GetFullRequestOriginStatByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) SetStatusCode(v int32) *GetFullRequestOriginStatByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) SetBody(v *GetFullRequestOriginStatByInstanceIdResponseBody) *GetFullRequestOriginStatByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) Validate() error {
	return dara.Validate(s)
}
