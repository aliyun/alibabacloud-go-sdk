// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeletePolarFsQuotaRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *DeletePolarFsQuotaRequest
	GetPolarFsInstanceId() *string
	SetQuotas(v []*DeletePolarFsQuotaRequestQuotas) *DeletePolarFsQuotaRequest
	GetQuotas() []*DeletePolarFsQuotaRequestQuotas
}

type DeletePolarFsQuotaRequest struct {
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// This parameter is required.
	Quotas []*DeletePolarFsQuotaRequestQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
}

func (s DeletePolarFsQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsQuotaRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarFsQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarFsQuotaRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsQuotaRequest) GetQuotas() []*DeletePolarFsQuotaRequestQuotas {
	return s.Quotas
}

func (s *DeletePolarFsQuotaRequest) SetDBClusterId(v string) *DeletePolarFsQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarFsQuotaRequest) SetPolarFsInstanceId(v string) *DeletePolarFsQuotaRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsQuotaRequest) SetQuotas(v []*DeletePolarFsQuotaRequestQuotas) *DeletePolarFsQuotaRequest {
	s.Quotas = v
	return s
}

func (s *DeletePolarFsQuotaRequest) Validate() error {
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

type DeletePolarFsQuotaRequestQuotas struct {
	// This parameter is required.
	//
	// example:
	//
	// 73
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sftest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeletePolarFsQuotaRequestQuotas) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsQuotaRequestQuotas) GoString() string {
	return s.String()
}

func (s *DeletePolarFsQuotaRequestQuotas) GetId() *string {
	return s.Id
}

func (s *DeletePolarFsQuotaRequestQuotas) GetName() *string {
	return s.Name
}

func (s *DeletePolarFsQuotaRequestQuotas) SetId(v string) *DeletePolarFsQuotaRequestQuotas {
	s.Id = &v
	return s
}

func (s *DeletePolarFsQuotaRequestQuotas) SetName(v string) *DeletePolarFsQuotaRequestQuotas {
	s.Name = &v
	return s
}

func (s *DeletePolarFsQuotaRequestQuotas) Validate() error {
	return dara.Validate(s)
}
