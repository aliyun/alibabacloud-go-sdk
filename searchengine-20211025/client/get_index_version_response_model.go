// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexVersionResponseBody) *GetIndexVersionResponse
	GetBody() *GetIndexVersionResponseBody
}

type GetIndexVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexVersionResponse) GetBody() *GetIndexVersionResponseBody {
	return s.Body
}

func (s *GetIndexVersionResponse) SetHeaders(v map[string]*string) *GetIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *GetIndexVersionResponse) SetStatusCode(v int32) *GetIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexVersionResponse) SetBody(v *GetIndexVersionResponseBody) *GetIndexVersionResponse {
	s.Body = v
	return s
}

func (s *GetIndexVersionResponse) Validate() error {
	return dara.Validate(s)
}
