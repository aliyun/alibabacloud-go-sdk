// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetDetailByUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetAssetDetailByUuidRequest
	GetLang() *string
	SetSourceIp(v string) *GetAssetDetailByUuidRequest
	GetSourceIp() *string
	SetUuid(v string) *GetAssetDetailByUuidRequest
	GetUuid() *string
}

type GetAssetDetailByUuidRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains this value.
	//
	// example:
	//
	// 120.245.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the asset that you want to query.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9e6cad93-a379-46fd-a701-9bbf02f4****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAssetDetailByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetDetailByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetAssetDetailByUuidRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAssetDetailByUuidRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetAssetDetailByUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAssetDetailByUuidRequest) SetLang(v string) *GetAssetDetailByUuidRequest {
	s.Lang = &v
	return s
}

func (s *GetAssetDetailByUuidRequest) SetSourceIp(v string) *GetAssetDetailByUuidRequest {
	s.SourceIp = &v
	return s
}

func (s *GetAssetDetailByUuidRequest) SetUuid(v string) *GetAssetDetailByUuidRequest {
	s.Uuid = &v
	return s
}

func (s *GetAssetDetailByUuidRequest) Validate() error {
	return dara.Validate(s)
}
