// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubPartnerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubPartnerListResponse
	GetStatusCode() *int32
	SetBody(v *GetSubPartnerListResponseBody) *GetSubPartnerListResponse
	GetBody() *GetSubPartnerListResponseBody
}

type GetSubPartnerListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubPartnerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubPartnerListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerListResponse) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubPartnerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubPartnerListResponse) GetBody() *GetSubPartnerListResponseBody {
	return s.Body
}

func (s *GetSubPartnerListResponse) SetHeaders(v map[string]*string) *GetSubPartnerListResponse {
	s.Headers = v
	return s
}

func (s *GetSubPartnerListResponse) SetStatusCode(v int32) *GetSubPartnerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubPartnerListResponse) SetBody(v *GetSubPartnerListResponseBody) *GetSubPartnerListResponse {
	s.Body = v
	return s
}

func (s *GetSubPartnerListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
