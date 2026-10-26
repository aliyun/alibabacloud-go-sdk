// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourcePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDataSourcePropertiesRequest
	GetDescription() *string
	SetDsId(v string) *UpdateDataSourcePropertiesRequest
	GetDsId() *string
	SetQuickBIOwnerId(v string) *UpdateDataSourcePropertiesRequest
	GetQuickBIOwnerId() *string
	SetShowName(v string) *UpdateDataSourcePropertiesRequest
	GetShowName() *string
}

type UpdateDataSourcePropertiesRequest struct {
	// example:
	//
	// ssss
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// example:
	//
	// sss
	QuickBIOwnerId *string `json:"QuickBIOwnerId,omitempty" xml:"QuickBIOwnerId,omitempty"`
	// example:
	//
	// wls
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s UpdateDataSourcePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourcePropertiesRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourcePropertiesRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataSourcePropertiesRequest) GetDsId() *string {
	return s.DsId
}

func (s *UpdateDataSourcePropertiesRequest) GetQuickBIOwnerId() *string {
	return s.QuickBIOwnerId
}

func (s *UpdateDataSourcePropertiesRequest) GetShowName() *string {
	return s.ShowName
}

func (s *UpdateDataSourcePropertiesRequest) SetDescription(v string) *UpdateDataSourcePropertiesRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourcePropertiesRequest) SetDsId(v string) *UpdateDataSourcePropertiesRequest {
	s.DsId = &v
	return s
}

func (s *UpdateDataSourcePropertiesRequest) SetQuickBIOwnerId(v string) *UpdateDataSourcePropertiesRequest {
	s.QuickBIOwnerId = &v
	return s
}

func (s *UpdateDataSourcePropertiesRequest) SetShowName(v string) *UpdateDataSourcePropertiesRequest {
	s.ShowName = &v
	return s
}

func (s *UpdateDataSourcePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
