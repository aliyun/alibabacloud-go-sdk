// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemoteADBDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRemoteADBDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRemoteADBDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRemoteADBDataSourceResponseBody) *ModifyRemoteADBDataSourceResponse
	GetBody() *ModifyRemoteADBDataSourceResponseBody
}

type ModifyRemoteADBDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRemoteADBDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRemoteADBDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemoteADBDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemoteADBDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRemoteADBDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRemoteADBDataSourceResponse) GetBody() *ModifyRemoteADBDataSourceResponseBody {
	return s.Body
}

func (s *ModifyRemoteADBDataSourceResponse) SetHeaders(v map[string]*string) *ModifyRemoteADBDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyRemoteADBDataSourceResponse) SetStatusCode(v int32) *ModifyRemoteADBDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponse) SetBody(v *ModifyRemoteADBDataSourceResponseBody) *ModifyRemoteADBDataSourceResponse {
	s.Body = v
	return s
}

func (s *ModifyRemoteADBDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
