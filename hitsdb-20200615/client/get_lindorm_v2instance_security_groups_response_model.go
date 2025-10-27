// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2InstanceSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2InstanceSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2InstanceSecurityGroupsResponseBody) *GetLindormV2InstanceSecurityGroupsResponse
	GetBody() *GetLindormV2InstanceSecurityGroupsResponseBody
}

type GetLindormV2InstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) GetBody() *GetLindormV2InstanceSecurityGroupsResponseBody {
	return s.Body
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) SetStatusCode(v int32) *GetLindormV2InstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) SetBody(v *GetLindormV2InstanceSecurityGroupsResponseBody) *GetLindormV2InstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
