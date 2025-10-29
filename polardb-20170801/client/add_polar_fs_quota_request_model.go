// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarFsQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddPolarFsQuotaRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *AddPolarFsQuotaRequest
	GetPolarFsInstanceId() *string
	SetQuotas(v []*AddPolarFsQuotaRequestQuotas) *AddPolarFsQuotaRequest
	GetQuotas() []*AddPolarFsQuotaRequestQuotas
}

type AddPolarFsQuotaRequest struct {
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// This parameter is required.
	Quotas []*AddPolarFsQuotaRequestQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
}

func (s AddPolarFsQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsQuotaRequest) GoString() string {
	return s.String()
}

func (s *AddPolarFsQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddPolarFsQuotaRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *AddPolarFsQuotaRequest) GetQuotas() []*AddPolarFsQuotaRequestQuotas {
	return s.Quotas
}

func (s *AddPolarFsQuotaRequest) SetDBClusterId(v string) *AddPolarFsQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddPolarFsQuotaRequest) SetPolarFsInstanceId(v string) *AddPolarFsQuotaRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *AddPolarFsQuotaRequest) SetQuotas(v []*AddPolarFsQuotaRequestQuotas) *AddPolarFsQuotaRequest {
	s.Quotas = v
	return s
}

func (s *AddPolarFsQuotaRequest) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddPolarFsQuotaRequestQuotas struct {
	// example:
	//
	// 7200
	AccessTTL *int64 `json:"AccessTTL,omitempty" xml:"AccessTTL,omitempty"`
	// example:
	//
	// 7200
	ChangeTTL *int64 `json:"ChangeTTL,omitempty" xml:"ChangeTTL,omitempty"`
	// example:
	//
	// quota_policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// /a/*project*
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// 222
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /b/*project*
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s AddPolarFsQuotaRequestQuotas) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsQuotaRequestQuotas) GoString() string {
	return s.String()
}

func (s *AddPolarFsQuotaRequestQuotas) GetAccessTTL() *int64 {
	return s.AccessTTL
}

func (s *AddPolarFsQuotaRequestQuotas) GetChangeTTL() *int64 {
	return s.ChangeTTL
}

func (s *AddPolarFsQuotaRequestQuotas) GetDescription() *string {
	return s.Description
}

func (s *AddPolarFsQuotaRequestQuotas) GetEnabled() *bool {
	return s.Enabled
}

func (s *AddPolarFsQuotaRequestQuotas) GetExclude() *string {
	return s.Exclude
}

func (s *AddPolarFsQuotaRequestQuotas) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *AddPolarFsQuotaRequestQuotas) GetInclude() *string {
	return s.Include
}

func (s *AddPolarFsQuotaRequestQuotas) GetName() *string {
	return s.Name
}

func (s *AddPolarFsQuotaRequestQuotas) GetPriority() *int32 {
	return s.Priority
}

func (s *AddPolarFsQuotaRequestQuotas) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *AddPolarFsQuotaRequestQuotas) SetAccessTTL(v int64) *AddPolarFsQuotaRequestQuotas {
	s.AccessTTL = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetChangeTTL(v int64) *AddPolarFsQuotaRequestQuotas {
	s.ChangeTTL = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetDescription(v string) *AddPolarFsQuotaRequestQuotas {
	s.Description = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetEnabled(v bool) *AddPolarFsQuotaRequestQuotas {
	s.Enabled = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetExclude(v string) *AddPolarFsQuotaRequestQuotas {
	s.Exclude = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetFileCountLimit(v int64) *AddPolarFsQuotaRequestQuotas {
	s.FileCountLimit = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetInclude(v string) *AddPolarFsQuotaRequestQuotas {
	s.Include = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetName(v string) *AddPolarFsQuotaRequestQuotas {
	s.Name = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetPriority(v int32) *AddPolarFsQuotaRequestQuotas {
	s.Priority = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) SetSizeLimit(v int64) *AddPolarFsQuotaRequestQuotas {
	s.SizeLimit = &v
	return s
}

func (s *AddPolarFsQuotaRequestQuotas) Validate() error {
	return dara.Validate(s)
}
