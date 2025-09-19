// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAssetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateAssetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateAssetGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateAssetGroupResponseBody) *CreateOrUpdateAssetGroupResponse
	GetBody() *CreateOrUpdateAssetGroupResponseBody
}

type CreateOrUpdateAssetGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateAssetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateAssetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateAssetGroupResponse) GetBody() *CreateOrUpdateAssetGroupResponseBody {
	return s.Body
}

func (s *CreateOrUpdateAssetGroupResponse) SetHeaders(v map[string]*string) *CreateOrUpdateAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateAssetGroupResponse) SetStatusCode(v int32) *CreateOrUpdateAssetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateAssetGroupResponse) SetBody(v *CreateOrUpdateAssetGroupResponseBody) *CreateOrUpdateAssetGroupResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateAssetGroupResponse) Validate() error {
	return dara.Validate(s)
}
