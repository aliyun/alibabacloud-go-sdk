// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectManagerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveProjectManagerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveProjectManagerResponse
	GetStatusCode() *int32
	SetBody(v *RemoveProjectManagerResponseBody) *RemoveProjectManagerResponse
	GetBody() *RemoveProjectManagerResponseBody
}

type RemoveProjectManagerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveProjectManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveProjectManagerResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerResponse) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveProjectManagerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveProjectManagerResponse) GetBody() *RemoveProjectManagerResponseBody {
	return s.Body
}

func (s *RemoveProjectManagerResponse) SetHeaders(v map[string]*string) *RemoveProjectManagerResponse {
	s.Headers = v
	return s
}

func (s *RemoveProjectManagerResponse) SetStatusCode(v int32) *RemoveProjectManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveProjectManagerResponse) SetBody(v *RemoveProjectManagerResponseBody) *RemoveProjectManagerResponse {
	s.Body = v
	return s
}

func (s *RemoveProjectManagerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
