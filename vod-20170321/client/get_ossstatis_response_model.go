// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOSSStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOSSStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetOSSStatisResponseBody) *GetOSSStatisResponse
	GetBody() *GetOSSStatisResponseBody
}

type GetOSSStatisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOSSStatisResponse) GoString() string {
	return s.String()
}

func (s *GetOSSStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOSSStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOSSStatisResponse) GetBody() *GetOSSStatisResponseBody {
	return s.Body
}

func (s *GetOSSStatisResponse) SetHeaders(v map[string]*string) *GetOSSStatisResponse {
	s.Headers = v
	return s
}

func (s *GetOSSStatisResponse) SetStatusCode(v int32) *GetOSSStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSStatisResponse) SetBody(v *GetOSSStatisResponseBody) *GetOSSStatisResponse {
	s.Body = v
	return s
}

func (s *GetOSSStatisResponse) Validate() error {
	return dara.Validate(s)
}
