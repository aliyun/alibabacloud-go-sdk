// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGtmResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveGtmResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveGtmResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *MoveGtmResourceGroupResponseBody) *MoveGtmResourceGroupResponse
	GetBody() *MoveGtmResourceGroupResponseBody
}

type MoveGtmResourceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveGtmResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveGtmResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveGtmResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveGtmResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveGtmResourceGroupResponse) GetBody() *MoveGtmResourceGroupResponseBody {
	return s.Body
}

func (s *MoveGtmResourceGroupResponse) SetHeaders(v map[string]*string) *MoveGtmResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveGtmResourceGroupResponse) SetStatusCode(v int32) *MoveGtmResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveGtmResourceGroupResponse) SetBody(v *MoveGtmResourceGroupResponseBody) *MoveGtmResourceGroupResponse {
	s.Body = v
	return s
}

func (s *MoveGtmResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
