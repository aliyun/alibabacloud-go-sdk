// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRdAccountDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *RdAccountDTO
	GetAccountType() *string
	SetChecked(v bool) *RdAccountDTO
	GetChecked() *bool
	SetDisplayName(v string) *RdAccountDTO
	GetDisplayName() *string
	SetId(v int64) *RdAccountDTO
	GetId() *int64
	SetName(v string) *RdAccountDTO
	GetName() *string
	SetTags(v []*RdAccountDTOTags) *RdAccountDTO
	GetTags() []*RdAccountDTOTags
}

type RdAccountDTO struct {
	AccountType *string             `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Checked     *bool               `json:"Checked,omitempty" xml:"Checked,omitempty"`
	DisplayName *string             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *int64              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags        []*RdAccountDTOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RdAccountDTO) String() string {
	return dara.Prettify(s)
}

func (s RdAccountDTO) GoString() string {
	return s.String()
}

func (s *RdAccountDTO) GetAccountType() *string {
	return s.AccountType
}

func (s *RdAccountDTO) GetChecked() *bool {
	return s.Checked
}

func (s *RdAccountDTO) GetDisplayName() *string {
	return s.DisplayName
}

func (s *RdAccountDTO) GetId() *int64 {
	return s.Id
}

func (s *RdAccountDTO) GetName() *string {
	return s.Name
}

func (s *RdAccountDTO) GetTags() []*RdAccountDTOTags {
	return s.Tags
}

func (s *RdAccountDTO) SetAccountType(v string) *RdAccountDTO {
	s.AccountType = &v
	return s
}

func (s *RdAccountDTO) SetChecked(v bool) *RdAccountDTO {
	s.Checked = &v
	return s
}

func (s *RdAccountDTO) SetDisplayName(v string) *RdAccountDTO {
	s.DisplayName = &v
	return s
}

func (s *RdAccountDTO) SetId(v int64) *RdAccountDTO {
	s.Id = &v
	return s
}

func (s *RdAccountDTO) SetName(v string) *RdAccountDTO {
	s.Name = &v
	return s
}

func (s *RdAccountDTO) SetTags(v []*RdAccountDTOTags) *RdAccountDTO {
	s.Tags = v
	return s
}

func (s *RdAccountDTO) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RdAccountDTOTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s RdAccountDTOTags) String() string {
	return dara.Prettify(s)
}

func (s RdAccountDTOTags) GoString() string {
	return s.String()
}

func (s *RdAccountDTOTags) GetTagKey() *string {
	return s.TagKey
}

func (s *RdAccountDTOTags) GetTagValue() *string {
	return s.TagValue
}

func (s *RdAccountDTOTags) SetTagKey(v string) *RdAccountDTOTags {
	s.TagKey = &v
	return s
}

func (s *RdAccountDTOTags) SetTagValue(v string) *RdAccountDTOTags {
	s.TagValue = &v
	return s
}

func (s *RdAccountDTOTags) Validate() error {
	return dara.Validate(s)
}
