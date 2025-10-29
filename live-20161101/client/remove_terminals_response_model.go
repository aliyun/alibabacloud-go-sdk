// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTerminalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTerminalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTerminalsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTerminalsResponseBody) *RemoveTerminalsResponse
	GetBody() *RemoveTerminalsResponseBody
}

type RemoveTerminalsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTerminalsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTerminalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTerminalsResponse) GetBody() *RemoveTerminalsResponseBody {
	return s.Body
}

func (s *RemoveTerminalsResponse) SetHeaders(v map[string]*string) *RemoveTerminalsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTerminalsResponse) SetStatusCode(v int32) *RemoveTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTerminalsResponse) SetBody(v *RemoveTerminalsResponseBody) *RemoveTerminalsResponse {
	s.Body = v
	return s
}

func (s *RemoveTerminalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
