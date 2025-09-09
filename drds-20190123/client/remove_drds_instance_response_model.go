// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDrdsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDrdsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDrdsInstanceResponseBody) *RemoveDrdsInstanceResponse
	GetBody() *RemoveDrdsInstanceResponseBody
}

type RemoveDrdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDrdsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDrdsInstanceResponse) GetBody() *RemoveDrdsInstanceResponseBody {
	return s.Body
}

func (s *RemoveDrdsInstanceResponse) SetHeaders(v map[string]*string) *RemoveDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetStatusCode(v int32) *RemoveDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetBody(v *RemoveDrdsInstanceResponseBody) *RemoveDrdsInstanceResponse {
	s.Body = v
	return s
}

func (s *RemoveDrdsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
