// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddressResp interface {
	dara.Model
	String() string
	GoString() string
	SetAddressDetail(v *AddressDetail) *GetAddressResp
	GetAddressDetail() *AddressDetail
	SetCreateTime(v string) *GetAddressResp
	GetCreateTime() *string
	SetModifyTime(v string) *GetAddressResp
	GetModifyTime() *string
	SetName(v string) *GetAddressResp
	GetName() *string
	SetOwner(v string) *GetAddressResp
	GetOwner() *string
	SetStatus(v string) *GetAddressResp
	GetStatus() *string
	SetTags(v string) *GetAddressResp
	GetTags() *string
	SetVerifyResult(v *VerifyResp) *GetAddressResp
	GetVerifyResult() *VerifyResp
	SetVerifyTime(v string) *GetAddressResp
	GetVerifyTime() *string
	SetVersion(v string) *GetAddressResp
	GetVersion() *string
}

type GetAddressResp struct {
	AddressDetail *AddressDetail `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// <your-address-name>
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1***90**87***53*
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// avaliable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags         *string     `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VerifyResult *VerifyResp `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	VerifyTime *string `json:"VerifyTime,omitempty" xml:"VerifyTime,omitempty"`
	// example:
	//
	// ****sf-****-0078-****-drfg****df1334
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAddressResp) String() string {
	return dara.Prettify(s)
}

func (s GetAddressResp) GoString() string {
	return s.String()
}

func (s *GetAddressResp) GetAddressDetail() *AddressDetail {
	return s.AddressDetail
}

func (s *GetAddressResp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAddressResp) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAddressResp) GetName() *string {
	return s.Name
}

func (s *GetAddressResp) GetOwner() *string {
	return s.Owner
}

func (s *GetAddressResp) GetStatus() *string {
	return s.Status
}

func (s *GetAddressResp) GetTags() *string {
	return s.Tags
}

func (s *GetAddressResp) GetVerifyResult() *VerifyResp {
	return s.VerifyResult
}

func (s *GetAddressResp) GetVerifyTime() *string {
	return s.VerifyTime
}

func (s *GetAddressResp) GetVersion() *string {
	return s.Version
}

func (s *GetAddressResp) SetAddressDetail(v *AddressDetail) *GetAddressResp {
	s.AddressDetail = v
	return s
}

func (s *GetAddressResp) SetCreateTime(v string) *GetAddressResp {
	s.CreateTime = &v
	return s
}

func (s *GetAddressResp) SetModifyTime(v string) *GetAddressResp {
	s.ModifyTime = &v
	return s
}

func (s *GetAddressResp) SetName(v string) *GetAddressResp {
	s.Name = &v
	return s
}

func (s *GetAddressResp) SetOwner(v string) *GetAddressResp {
	s.Owner = &v
	return s
}

func (s *GetAddressResp) SetStatus(v string) *GetAddressResp {
	s.Status = &v
	return s
}

func (s *GetAddressResp) SetTags(v string) *GetAddressResp {
	s.Tags = &v
	return s
}

func (s *GetAddressResp) SetVerifyResult(v *VerifyResp) *GetAddressResp {
	s.VerifyResult = v
	return s
}

func (s *GetAddressResp) SetVerifyTime(v string) *GetAddressResp {
	s.VerifyTime = &v
	return s
}

func (s *GetAddressResp) SetVersion(v string) *GetAddressResp {
	s.Version = &v
	return s
}

func (s *GetAddressResp) Validate() error {
	if s.AddressDetail != nil {
		if err := s.AddressDetail.Validate(); err != nil {
			return err
		}
	}
	if s.VerifyResult != nil {
		if err := s.VerifyResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
