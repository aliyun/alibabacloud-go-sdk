// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLabelTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLabelTableResponse
	GetStatusCode() *int32
	SetBody(v *GetLabelTableResponseBody) *GetLabelTableResponse
	GetBody() *GetLabelTableResponseBody
}

type GetLabelTableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLabelTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLabelTableResponse) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLabelTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLabelTableResponse) GetBody() *GetLabelTableResponseBody {
	return s.Body
}

func (s *GetLabelTableResponse) SetHeaders(v map[string]*string) *GetLabelTableResponse {
	s.Headers = v
	return s
}

func (s *GetLabelTableResponse) SetStatusCode(v int32) *GetLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelTableResponse) SetBody(v *GetLabelTableResponseBody) *GetLabelTableResponse {
	s.Body = v
	return s
}

func (s *GetLabelTableResponse) Validate() error {
	return dara.Validate(s)
}
