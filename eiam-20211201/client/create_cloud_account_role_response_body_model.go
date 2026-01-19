// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountRoleId(v string) *CreateCloudAccountRoleResponseBody
	GetCloudAccountRoleId() *string
	SetRequestId(v string) *CreateCloudAccountRoleResponseBody
	GetRequestId() *string
}

type CreateCloudAccountRoleResponseBody struct {
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRoleResponseBody) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *CreateCloudAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudAccountRoleResponseBody) SetCloudAccountRoleId(v string) *CreateCloudAccountRoleResponseBody {
	s.CloudAccountRoleId = &v
	return s
}

func (s *CreateCloudAccountRoleResponseBody) SetRequestId(v string) *CreateCloudAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudAccountRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
