// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestSampleByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFullRequestSampleByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFullRequestSampleByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetFullRequestSampleByInstanceIdResponseBody) *GetFullRequestSampleByInstanceIdResponse
	GetBody() *GetFullRequestSampleByInstanceIdResponseBody
}

type GetFullRequestSampleByInstanceIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFullRequestSampleByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFullRequestSampleByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestSampleByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetFullRequestSampleByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFullRequestSampleByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFullRequestSampleByInstanceIdResponse) GetBody() *GetFullRequestSampleByInstanceIdResponseBody {
	return s.Body
}

func (s *GetFullRequestSampleByInstanceIdResponse) SetHeaders(v map[string]*string) *GetFullRequestSampleByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponse) SetStatusCode(v int32) *GetFullRequestSampleByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponse) SetBody(v *GetFullRequestSampleByInstanceIdResponseBody) *GetFullRequestSampleByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponse) Validate() error {
	return dara.Validate(s)
}
