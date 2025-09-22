// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAreaElecConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAreaElecConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAreaElecConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetAreaElecConstituteResponseBody) *GetAreaElecConstituteResponse
	GetBody() *GetAreaElecConstituteResponseBody
}

type GetAreaElecConstituteResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAreaElecConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAreaElecConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAreaElecConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAreaElecConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAreaElecConstituteResponse) GetBody() *GetAreaElecConstituteResponseBody {
	return s.Body
}

func (s *GetAreaElecConstituteResponse) SetHeaders(v map[string]*string) *GetAreaElecConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetAreaElecConstituteResponse) SetStatusCode(v int32) *GetAreaElecConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAreaElecConstituteResponse) SetBody(v *GetAreaElecConstituteResponseBody) *GetAreaElecConstituteResponse {
	s.Body = v
	return s
}

func (s *GetAreaElecConstituteResponse) Validate() error {
	return dara.Validate(s)
}
