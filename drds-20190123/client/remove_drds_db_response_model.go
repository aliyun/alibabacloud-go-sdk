// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDrdsDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDrdsDbResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDrdsDbResponseBody) *RemoveDrdsDbResponse
	GetBody() *RemoveDrdsDbResponseBody
}

type RemoveDrdsDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsDbResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDrdsDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDrdsDbResponse) GetBody() *RemoveDrdsDbResponseBody {
	return s.Body
}

func (s *RemoveDrdsDbResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbResponse) SetStatusCode(v int32) *RemoveDrdsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsDbResponse) SetBody(v *RemoveDrdsDbResponseBody) *RemoveDrdsDbResponse {
	s.Body = v
	return s
}

func (s *RemoveDrdsDbResponse) Validate() error {
	return dara.Validate(s)
}
