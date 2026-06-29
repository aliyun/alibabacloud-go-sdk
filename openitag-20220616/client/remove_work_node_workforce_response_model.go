// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkNodeWorkforceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveWorkNodeWorkforceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveWorkNodeWorkforceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveWorkNodeWorkforceResponseBody) *RemoveWorkNodeWorkforceResponse
	GetBody() *RemoveWorkNodeWorkforceResponseBody
}

type RemoveWorkNodeWorkforceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWorkNodeWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveWorkNodeWorkforceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkNodeWorkforceResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveWorkNodeWorkforceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveWorkNodeWorkforceResponse) GetBody() *RemoveWorkNodeWorkforceResponseBody {
	return s.Body
}

func (s *RemoveWorkNodeWorkforceResponse) SetHeaders(v map[string]*string) *RemoveWorkNodeWorkforceResponse {
	s.Headers = v
	return s
}

func (s *RemoveWorkNodeWorkforceResponse) SetStatusCode(v int32) *RemoveWorkNodeWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponse) SetBody(v *RemoveWorkNodeWorkforceResponseBody) *RemoveWorkNodeWorkforceResponse {
	s.Body = v
	return s
}

func (s *RemoveWorkNodeWorkforceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
