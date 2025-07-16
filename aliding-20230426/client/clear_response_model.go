// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearResponse
	GetStatusCode() *int32
	SetBody(v *ClearResponseBody) *ClearResponse
	GetBody() *ClearResponseBody
}

type ClearResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearResponse) GoString() string {
	return s.String()
}

func (s *ClearResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearResponse) GetBody() *ClearResponseBody {
	return s.Body
}

func (s *ClearResponse) SetHeaders(v map[string]*string) *ClearResponse {
	s.Headers = v
	return s
}

func (s *ClearResponse) SetStatusCode(v int32) *ClearResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearResponse) SetBody(v *ClearResponseBody) *ClearResponse {
	s.Body = v
	return s
}

func (s *ClearResponse) Validate() error {
	return dara.Validate(s)
}
