// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServiceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpServiceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpServiceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceGroupPageResult) *ListPdpServiceGroupsResponse
	GetBody() *PdpServiceGroupPageResult
}

type ListPdpServiceGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceGroupPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpServiceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPdpServiceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpServiceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpServiceGroupsResponse) GetBody() *PdpServiceGroupPageResult {
	return s.Body
}

func (s *ListPdpServiceGroupsResponse) SetHeaders(v map[string]*string) *ListPdpServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPdpServiceGroupsResponse) SetStatusCode(v int32) *ListPdpServiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpServiceGroupsResponse) SetBody(v *PdpServiceGroupPageResult) *ListPdpServiceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListPdpServiceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
