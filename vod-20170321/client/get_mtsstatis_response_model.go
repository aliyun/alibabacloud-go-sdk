// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMTSStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMTSStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMTSStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetMTSStatisResponseBody) *GetMTSStatisResponse
	GetBody() *GetMTSStatisResponseBody
}

type GetMTSStatisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMTSStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMTSStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponse) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMTSStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMTSStatisResponse) GetBody() *GetMTSStatisResponseBody {
	return s.Body
}

func (s *GetMTSStatisResponse) SetHeaders(v map[string]*string) *GetMTSStatisResponse {
	s.Headers = v
	return s
}

func (s *GetMTSStatisResponse) SetStatusCode(v int32) *GetMTSStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMTSStatisResponse) SetBody(v *GetMTSStatisResponseBody) *GetMTSStatisResponse {
	s.Body = v
	return s
}

func (s *GetMTSStatisResponse) Validate() error {
	return dara.Validate(s)
}
