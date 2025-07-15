// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVpcAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVpcAccessResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVpcAccessResponseBody) *RemoveVpcAccessResponse
	GetBody() *RemoveVpcAccessResponseBody
}

type RemoveVpcAccessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVpcAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVpcAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessResponse) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVpcAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVpcAccessResponse) GetBody() *RemoveVpcAccessResponseBody {
	return s.Body
}

func (s *RemoveVpcAccessResponse) SetHeaders(v map[string]*string) *RemoveVpcAccessResponse {
	s.Headers = v
	return s
}

func (s *RemoveVpcAccessResponse) SetStatusCode(v int32) *RemoveVpcAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVpcAccessResponse) SetBody(v *RemoveVpcAccessResponseBody) *RemoveVpcAccessResponse {
	s.Body = v
	return s
}

func (s *RemoveVpcAccessResponse) Validate() error {
	return dara.Validate(s)
}
