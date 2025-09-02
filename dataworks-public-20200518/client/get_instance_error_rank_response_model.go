// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceErrorRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceErrorRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceErrorRankResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceErrorRankResponseBody) *GetInstanceErrorRankResponse
	GetBody() *GetInstanceErrorRankResponseBody
}

type GetInstanceErrorRankResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceErrorRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceErrorRankResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceErrorRankResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceErrorRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceErrorRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceErrorRankResponse) GetBody() *GetInstanceErrorRankResponseBody {
	return s.Body
}

func (s *GetInstanceErrorRankResponse) SetHeaders(v map[string]*string) *GetInstanceErrorRankResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceErrorRankResponse) SetStatusCode(v int32) *GetInstanceErrorRankResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceErrorRankResponse) SetBody(v *GetInstanceErrorRankResponseBody) *GetInstanceErrorRankResponse {
	s.Body = v
	return s
}

func (s *GetInstanceErrorRankResponse) Validate() error {
	return dara.Validate(s)
}
