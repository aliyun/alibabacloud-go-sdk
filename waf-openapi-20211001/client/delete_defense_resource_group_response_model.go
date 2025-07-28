// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDefenseResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDefenseResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDefenseResourceGroupResponseBody) *DeleteDefenseResourceGroupResponse
	GetBody() *DeleteDefenseResourceGroupResponseBody
}

type DeleteDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDefenseResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDefenseResourceGroupResponse) GetBody() *DeleteDefenseResourceGroupResponseBody {
	return s.Body
}

func (s *DeleteDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetStatusCode(v int32) *DeleteDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetBody(v *DeleteDefenseResourceGroupResponseBody) *DeleteDefenseResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
