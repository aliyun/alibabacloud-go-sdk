// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetElecTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetElecTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetElecTrendResponseBody) *GetElecTrendResponse
	GetBody() *GetElecTrendResponseBody
}

type GetElecTrendResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElecTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElecTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponse) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetElecTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetElecTrendResponse) GetBody() *GetElecTrendResponseBody {
	return s.Body
}

func (s *GetElecTrendResponse) SetHeaders(v map[string]*string) *GetElecTrendResponse {
	s.Headers = v
	return s
}

func (s *GetElecTrendResponse) SetStatusCode(v int32) *GetElecTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElecTrendResponse) SetBody(v *GetElecTrendResponseBody) *GetElecTrendResponse {
	s.Body = v
	return s
}

func (s *GetElecTrendResponse) Validate() error {
	return dara.Validate(s)
}
