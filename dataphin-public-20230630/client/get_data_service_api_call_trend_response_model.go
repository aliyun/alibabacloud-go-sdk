// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiCallTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiCallTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiCallTrendResponseBody) *GetDataServiceApiCallTrendResponse
	GetBody() *GetDataServiceApiCallTrendResponseBody
}

type GetDataServiceApiCallTrendResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiCallTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiCallTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiCallTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiCallTrendResponse) GetBody() *GetDataServiceApiCallTrendResponseBody {
	return s.Body
}

func (s *GetDataServiceApiCallTrendResponse) SetHeaders(v map[string]*string) *GetDataServiceApiCallTrendResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiCallTrendResponse) SetStatusCode(v int32) *GetDataServiceApiCallTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponse) SetBody(v *GetDataServiceApiCallTrendResponseBody) *GetDataServiceApiCallTrendResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiCallTrendResponse) Validate() error {
	return dara.Validate(s)
}
