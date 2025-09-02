// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaselineResponse
	GetStatusCode() *int32
	SetBody(v *GetBaselineResponseBody) *GetBaselineResponse
	GetBody() *GetBaselineResponseBody
}

type GetBaselineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaselineResponse) GetBody() *GetBaselineResponseBody {
	return s.Body
}

func (s *GetBaselineResponse) SetHeaders(v map[string]*string) *GetBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineResponse) SetStatusCode(v int32) *GetBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineResponse) SetBody(v *GetBaselineResponseBody) *GetBaselineResponse {
	s.Body = v
	return s
}

func (s *GetBaselineResponse) Validate() error {
	return dara.Validate(s)
}
