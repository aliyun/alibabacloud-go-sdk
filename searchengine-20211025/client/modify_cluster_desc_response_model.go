// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDescResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterDescResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterDescResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterDescResponseBody) *ModifyClusterDescResponse
	GetBody() *ModifyClusterDescResponseBody
}

type ModifyClusterDescResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterDescResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDescResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDescResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterDescResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterDescResponse) GetBody() *ModifyClusterDescResponseBody {
	return s.Body
}

func (s *ModifyClusterDescResponse) SetHeaders(v map[string]*string) *ModifyClusterDescResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDescResponse) SetStatusCode(v int32) *ModifyClusterDescResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterDescResponse) SetBody(v *ModifyClusterDescResponseBody) *ModifyClusterDescResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterDescResponse) Validate() error {
	return dara.Validate(s)
}
