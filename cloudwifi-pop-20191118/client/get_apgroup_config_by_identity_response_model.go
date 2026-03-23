// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupConfigByIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApgroupConfigByIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApgroupConfigByIdentityResponse
	GetStatusCode() *int32
	SetBody(v *GetApgroupConfigByIdentityResponseBody) *GetApgroupConfigByIdentityResponse
	GetBody() *GetApgroupConfigByIdentityResponseBody
}

type GetApgroupConfigByIdentityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApgroupConfigByIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApgroupConfigByIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupConfigByIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetApgroupConfigByIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApgroupConfigByIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApgroupConfigByIdentityResponse) GetBody() *GetApgroupConfigByIdentityResponseBody {
	return s.Body
}

func (s *GetApgroupConfigByIdentityResponse) SetHeaders(v map[string]*string) *GetApgroupConfigByIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetApgroupConfigByIdentityResponse) SetStatusCode(v int32) *GetApgroupConfigByIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApgroupConfigByIdentityResponse) SetBody(v *GetApgroupConfigByIdentityResponseBody) *GetApgroupConfigByIdentityResponse {
	s.Body = v
	return s
}

func (s *GetApgroupConfigByIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
