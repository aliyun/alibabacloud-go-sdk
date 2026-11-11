// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudPermissionTemplateResponseBody
	GetRequestId() *string
}

type DeleteCloudPermissionTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudPermissionTemplateResponseBody) SetRequestId(v string) *DeleteCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudPermissionTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
