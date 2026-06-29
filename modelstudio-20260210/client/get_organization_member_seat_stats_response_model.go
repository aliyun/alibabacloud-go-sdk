// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationMemberSeatStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrganizationMemberSeatStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrganizationMemberSeatStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetOrganizationMemberSeatStatsResponseBody) *GetOrganizationMemberSeatStatsResponse
	GetBody() *GetOrganizationMemberSeatStatsResponseBody
}

type GetOrganizationMemberSeatStatsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationMemberSeatStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationMemberSeatStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberSeatStatsResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberSeatStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrganizationMemberSeatStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrganizationMemberSeatStatsResponse) GetBody() *GetOrganizationMemberSeatStatsResponseBody {
	return s.Body
}

func (s *GetOrganizationMemberSeatStatsResponse) SetHeaders(v map[string]*string) *GetOrganizationMemberSeatStatsResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponse) SetStatusCode(v int32) *GetOrganizationMemberSeatStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponse) SetBody(v *GetOrganizationMemberSeatStatsResponseBody) *GetOrganizationMemberSeatStatsResponse {
	s.Body = v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
