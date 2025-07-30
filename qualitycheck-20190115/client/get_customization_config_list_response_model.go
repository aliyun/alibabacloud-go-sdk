// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizationConfigListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomizationConfigListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomizationConfigListResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomizationConfigListResponseBody) *GetCustomizationConfigListResponse
	GetBody() *GetCustomizationConfigListResponseBody
}

type GetCustomizationConfigListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomizationConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomizationConfigListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizationConfigListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomizationConfigListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomizationConfigListResponse) GetBody() *GetCustomizationConfigListResponseBody {
	return s.Body
}

func (s *GetCustomizationConfigListResponse) SetHeaders(v map[string]*string) *GetCustomizationConfigListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizationConfigListResponse) SetStatusCode(v int32) *GetCustomizationConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizationConfigListResponse) SetBody(v *GetCustomizationConfigListResponseBody) *GetCustomizationConfigListResponse {
	s.Body = v
	return s
}

func (s *GetCustomizationConfigListResponse) Validate() error {
	return dara.Validate(s)
}
