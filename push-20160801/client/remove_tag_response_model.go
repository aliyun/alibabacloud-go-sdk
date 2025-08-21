// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTagResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTagResponseBody) *RemoveTagResponse
	GetBody() *RemoveTagResponseBody
}

type RemoveTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTagResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTagResponse) GetBody() *RemoveTagResponseBody {
	return s.Body
}

func (s *RemoveTagResponse) SetHeaders(v map[string]*string) *RemoveTagResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagResponse) SetStatusCode(v int32) *RemoveTagResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagResponse) SetBody(v *RemoveTagResponseBody) *RemoveTagResponse {
	s.Body = v
	return s
}

func (s *RemoveTagResponse) Validate() error {
	return dara.Validate(s)
}
