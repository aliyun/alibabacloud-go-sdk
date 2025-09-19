// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetDataTrendResponseBody) *GetDataTrendResponse
	GetBody() *GetDataTrendResponseBody
}

type GetDataTrendResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrendResponse) GoString() string {
	return s.String()
}

func (s *GetDataTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataTrendResponse) GetBody() *GetDataTrendResponseBody {
	return s.Body
}

func (s *GetDataTrendResponse) SetHeaders(v map[string]*string) *GetDataTrendResponse {
	s.Headers = v
	return s
}

func (s *GetDataTrendResponse) SetStatusCode(v int32) *GetDataTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataTrendResponse) SetBody(v *GetDataTrendResponseBody) *GetDataTrendResponse {
	s.Body = v
	return s
}

func (s *GetDataTrendResponse) Validate() error {
	return dara.Validate(s)
}
