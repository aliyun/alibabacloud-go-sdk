// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSDGsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSDGsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSDGsResponseBody) *RemoveSDGsResponse
	GetBody() *RemoveSDGsResponseBody
}

type RemoveSDGsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSDGsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSDGsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsResponse) GoString() string {
	return s.String()
}

func (s *RemoveSDGsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSDGsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSDGsResponse) GetBody() *RemoveSDGsResponseBody {
	return s.Body
}

func (s *RemoveSDGsResponse) SetHeaders(v map[string]*string) *RemoveSDGsResponse {
	s.Headers = v
	return s
}

func (s *RemoveSDGsResponse) SetStatusCode(v int32) *RemoveSDGsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSDGsResponse) SetBody(v *RemoveSDGsResponseBody) *RemoveSDGsResponse {
	s.Body = v
	return s
}

func (s *RemoveSDGsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
