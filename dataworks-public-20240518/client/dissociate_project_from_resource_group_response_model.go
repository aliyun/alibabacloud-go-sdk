// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateProjectFromResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateProjectFromResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DissociateProjectFromResourceGroupResponseBody) *DissociateProjectFromResourceGroupResponse
	GetBody() *DissociateProjectFromResourceGroupResponseBody
}

type DissociateProjectFromResourceGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateProjectFromResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateProjectFromResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateProjectFromResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateProjectFromResourceGroupResponse) GetBody() *DissociateProjectFromResourceGroupResponseBody {
	return s.Body
}

func (s *DissociateProjectFromResourceGroupResponse) SetHeaders(v map[string]*string) *DissociateProjectFromResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DissociateProjectFromResourceGroupResponse) SetStatusCode(v int32) *DissociateProjectFromResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateProjectFromResourceGroupResponse) SetBody(v *DissociateProjectFromResourceGroupResponseBody) *DissociateProjectFromResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DissociateProjectFromResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
