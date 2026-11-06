// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudPermissionTemplateId(v string) *CreateCloudPermissionTemplateResponseBody
	GetCloudPermissionTemplateId() *string
	SetRequestId(v string) *CreateCloudPermissionTemplateResponseBody
	GetRequestId() *string
}

type CreateCloudPermissionTemplateResponseBody struct {
	// The ID of the permission template.
	//
	// example:
	//
	// cpt_01l8dagf2iablhigj75l1n99xxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudPermissionTemplateResponseBody) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *CreateCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudPermissionTemplateResponseBody) SetCloudPermissionTemplateId(v string) *CreateCloudPermissionTemplateResponseBody {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *CreateCloudPermissionTemplateResponseBody) SetRequestId(v string) *CreateCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudPermissionTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
