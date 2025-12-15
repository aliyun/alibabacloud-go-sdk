// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupAdminSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorAsAdmin(v bool) *GetResourceGroupAdminSettingResponseBody
	GetCreatorAsAdmin() *bool
	SetRequestId(v string) *GetResourceGroupAdminSettingResponseBody
	GetRequestId() *string
}

type GetResourceGroupAdminSettingResponseBody struct {
	// Indicates whether enable the Use Creator as Administrator feature.
	//
	// example:
	//
	// true
	CreatorAsAdmin *bool `json:"CreatorAsAdmin,omitempty" xml:"CreatorAsAdmin,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 776B05B3-A0B0-464B-A191-F4E1119A94B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceGroupAdminSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAdminSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAdminSettingResponseBody) GetCreatorAsAdmin() *bool {
	return s.CreatorAsAdmin
}

func (s *GetResourceGroupAdminSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupAdminSettingResponseBody) SetCreatorAsAdmin(v bool) *GetResourceGroupAdminSettingResponseBody {
	s.CreatorAsAdmin = &v
	return s
}

func (s *GetResourceGroupAdminSettingResponseBody) SetRequestId(v string) *GetResourceGroupAdminSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupAdminSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
