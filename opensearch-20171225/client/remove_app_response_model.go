// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAppResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAppResponseBody) *RemoveAppResponse
	GetBody() *RemoveAppResponseBody
}

type RemoveAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAppResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAppResponse) GetBody() *RemoveAppResponseBody {
	return s.Body
}

func (s *RemoveAppResponse) SetHeaders(v map[string]*string) *RemoveAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppResponse) SetStatusCode(v int32) *RemoveAppResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppResponse) SetBody(v *RemoveAppResponseBody) *RemoveAppResponse {
	s.Body = v
	return s
}

func (s *RemoveAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
