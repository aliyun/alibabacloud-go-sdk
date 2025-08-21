// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInstanceSDGResponseBody) *RemoveInstanceSDGResponse
	GetBody() *RemoveInstanceSDGResponseBody
}

type RemoveInstanceSDGResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInstanceSDGResponse) GetBody() *RemoveInstanceSDGResponseBody {
	return s.Body
}

func (s *RemoveInstanceSDGResponse) SetHeaders(v map[string]*string) *RemoveInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceSDGResponse) SetStatusCode(v int32) *RemoveInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceSDGResponse) SetBody(v *RemoveInstanceSDGResponseBody) *RemoveInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *RemoveInstanceSDGResponse) Validate() error {
	return dara.Validate(s)
}
