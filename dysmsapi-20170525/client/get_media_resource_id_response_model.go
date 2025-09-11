// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaResourceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaResourceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaResourceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaResourceIdResponseBody) *GetMediaResourceIdResponse
	GetBody() *GetMediaResourceIdResponseBody
}

type GetMediaResourceIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaResourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaResourceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResourceIdResponse) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaResourceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaResourceIdResponse) GetBody() *GetMediaResourceIdResponseBody {
	return s.Body
}

func (s *GetMediaResourceIdResponse) SetHeaders(v map[string]*string) *GetMediaResourceIdResponse {
	s.Headers = v
	return s
}

func (s *GetMediaResourceIdResponse) SetStatusCode(v int32) *GetMediaResourceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaResourceIdResponse) SetBody(v *GetMediaResourceIdResponseBody) *GetMediaResourceIdResponse {
	s.Body = v
	return s
}

func (s *GetMediaResourceIdResponse) Validate() error {
	return dara.Validate(s)
}
