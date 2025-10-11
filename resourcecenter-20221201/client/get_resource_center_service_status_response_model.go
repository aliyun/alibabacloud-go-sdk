// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCenterServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceCenterServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceCenterServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceCenterServiceStatusResponseBody) *GetResourceCenterServiceStatusResponse
	GetBody() *GetResourceCenterServiceStatusResponseBody
}

type GetResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCenterServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceCenterServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceCenterServiceStatusResponse) GetBody() *GetResourceCenterServiceStatusResponseBody {
	return s.Body
}

func (s *GetResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetBody(v *GetResourceCenterServiceStatusResponseBody) *GetResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
