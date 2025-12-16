// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteABTestGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteABTestGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteABTestGroupResponseBody) *DeleteABTestGroupResponse
	GetBody() *DeleteABTestGroupResponseBody
}

type DeleteABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABTestGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteABTestGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteABTestGroupResponse) GetBody() *DeleteABTestGroupResponseBody {
	return s.Body
}

func (s *DeleteABTestGroupResponse) SetHeaders(v map[string]*string) *DeleteABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestGroupResponse) SetStatusCode(v int32) *DeleteABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestGroupResponse) SetBody(v *DeleteABTestGroupResponseBody) *DeleteABTestGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteABTestGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
