// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WearOrgHonorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WearOrgHonorResponse
	GetStatusCode() *int32
	SetBody(v *WearOrgHonorResponseBody) *WearOrgHonorResponse
	GetBody() *WearOrgHonorResponseBody
}

type WearOrgHonorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WearOrgHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WearOrgHonorResponse) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorResponse) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WearOrgHonorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WearOrgHonorResponse) GetBody() *WearOrgHonorResponseBody {
	return s.Body
}

func (s *WearOrgHonorResponse) SetHeaders(v map[string]*string) *WearOrgHonorResponse {
	s.Headers = v
	return s
}

func (s *WearOrgHonorResponse) SetStatusCode(v int32) *WearOrgHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *WearOrgHonorResponse) SetBody(v *WearOrgHonorResponseBody) *WearOrgHonorResponse {
	s.Body = v
	return s
}

func (s *WearOrgHonorResponse) Validate() error {
	return dara.Validate(s)
}
