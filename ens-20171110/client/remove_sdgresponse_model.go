// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSDGResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSDGResponseBody) *RemoveSDGResponse
	GetBody() *RemoveSDGResponseBody
}

type RemoveSDGResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGResponse) GoString() string {
	return s.String()
}

func (s *RemoveSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSDGResponse) GetBody() *RemoveSDGResponseBody {
	return s.Body
}

func (s *RemoveSDGResponse) SetHeaders(v map[string]*string) *RemoveSDGResponse {
	s.Headers = v
	return s
}

func (s *RemoveSDGResponse) SetStatusCode(v int32) *RemoveSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSDGResponse) SetBody(v *RemoveSDGResponseBody) *RemoveSDGResponse {
	s.Body = v
	return s
}

func (s *RemoveSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
