// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeOnBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeOnBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeOnBaselineResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeOnBaselineResponseBody) *GetNodeOnBaselineResponse
	GetBody() *GetNodeOnBaselineResponseBody
}

type GetNodeOnBaselineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeOnBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeOnBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOnBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetNodeOnBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeOnBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeOnBaselineResponse) GetBody() *GetNodeOnBaselineResponseBody {
	return s.Body
}

func (s *GetNodeOnBaselineResponse) SetHeaders(v map[string]*string) *GetNodeOnBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetNodeOnBaselineResponse) SetStatusCode(v int32) *GetNodeOnBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeOnBaselineResponse) SetBody(v *GetNodeOnBaselineResponseBody) *GetNodeOnBaselineResponse {
	s.Body = v
	return s
}

func (s *GetNodeOnBaselineResponse) Validate() error {
	return dara.Validate(s)
}
