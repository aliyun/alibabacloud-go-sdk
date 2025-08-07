// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterServerlessConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterServerlessConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterServerlessConfResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterServerlessConfResponseBody) *ModifyDBClusterServerlessConfResponse
	GetBody() *ModifyDBClusterServerlessConfResponseBody
}

type ModifyDBClusterServerlessConfResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterServerlessConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterServerlessConfResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterServerlessConfResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterServerlessConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterServerlessConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterServerlessConfResponse) GetBody() *ModifyDBClusterServerlessConfResponseBody {
	return s.Body
}

func (s *ModifyDBClusterServerlessConfResponse) SetHeaders(v map[string]*string) *ModifyDBClusterServerlessConfResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterServerlessConfResponse) SetStatusCode(v int32) *ModifyDBClusterServerlessConfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterServerlessConfResponse) SetBody(v *ModifyDBClusterServerlessConfResponseBody) *ModifyDBClusterServerlessConfResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterServerlessConfResponse) Validate() error {
	return dara.Validate(s)
}
