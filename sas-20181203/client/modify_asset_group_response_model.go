// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAssetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAssetGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAssetGroupResponseBody) *ModifyAssetGroupResponse
	GetBody() *ModifyAssetGroupResponseBody
}

type ModifyAssetGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAssetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAssetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAssetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAssetGroupResponse) GetBody() *ModifyAssetGroupResponseBody {
	return s.Body
}

func (s *ModifyAssetGroupResponse) SetHeaders(v map[string]*string) *ModifyAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAssetGroupResponse) SetStatusCode(v int32) *ModifyAssetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAssetGroupResponse) SetBody(v *ModifyAssetGroupResponseBody) *ModifyAssetGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyAssetGroupResponse) Validate() error {
	return dara.Validate(s)
}
