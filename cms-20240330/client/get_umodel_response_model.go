// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUmodelResponse
	GetStatusCode() *int32
	SetBody(v *GetUmodelResponseBody) *GetUmodelResponse
	GetBody() *GetUmodelResponseBody
}

type GetUmodelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelResponse) GoString() string {
	return s.String()
}

func (s *GetUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUmodelResponse) GetBody() *GetUmodelResponseBody {
	return s.Body
}

func (s *GetUmodelResponse) SetHeaders(v map[string]*string) *GetUmodelResponse {
	s.Headers = v
	return s
}

func (s *GetUmodelResponse) SetStatusCode(v int32) *GetUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUmodelResponse) SetBody(v *GetUmodelResponseBody) *GetUmodelResponse {
	s.Body = v
	return s
}

func (s *GetUmodelResponse) Validate() error {
	return dara.Validate(s)
}
