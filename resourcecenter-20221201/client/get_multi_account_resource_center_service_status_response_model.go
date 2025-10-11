// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCenterServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiAccountResourceCenterServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiAccountResourceCenterServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiAccountResourceCenterServiceStatusResponseBody) *GetMultiAccountResourceCenterServiceStatusResponse
	GetBody() *GetMultiAccountResourceCenterServiceStatusResponseBody
}

type GetMultiAccountResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) GetBody() *GetMultiAccountResourceCenterServiceStatusResponseBody {
	return s.Body
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetBody(v *GetMultiAccountResourceCenterServiceStatusResponseBody) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
