// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSFlowStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOSSFlowStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOSSFlowStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetOSSFlowStatisResponseBody) *GetOSSFlowStatisResponse
	GetBody() *GetOSSFlowStatisResponseBody
}

type GetOSSFlowStatisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSFlowStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSFlowStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOSSFlowStatisResponse) GoString() string {
	return s.String()
}

func (s *GetOSSFlowStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOSSFlowStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOSSFlowStatisResponse) GetBody() *GetOSSFlowStatisResponseBody {
	return s.Body
}

func (s *GetOSSFlowStatisResponse) SetHeaders(v map[string]*string) *GetOSSFlowStatisResponse {
	s.Headers = v
	return s
}

func (s *GetOSSFlowStatisResponse) SetStatusCode(v int32) *GetOSSFlowStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSFlowStatisResponse) SetBody(v *GetOSSFlowStatisResponseBody) *GetOSSFlowStatisResponse {
	s.Body = v
	return s
}

func (s *GetOSSFlowStatisResponse) Validate() error {
	return dara.Validate(s)
}
