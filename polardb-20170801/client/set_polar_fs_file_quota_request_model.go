// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolarFsFileQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SetPolarFsFileQuotaRequest
	GetDBClusterId() *string
	SetFilePathQuotas(v []*SetPolarFsFileQuotaRequestFilePathQuotas) *SetPolarFsFileQuotaRequest
	GetFilePathQuotas() []*SetPolarFsFileQuotaRequestFilePathQuotas
	SetPolarFsInstanceId(v string) *SetPolarFsFileQuotaRequest
	GetPolarFsInstanceId() *string
}

type SetPolarFsFileQuotaRequest struct {
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	FilePathQuotas []*SetPolarFsFileQuotaRequestFilePathQuotas `json:"FilePathQuotas,omitempty" xml:"FilePathQuotas,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s SetPolarFsFileQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolarFsFileQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetPolarFsFileQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SetPolarFsFileQuotaRequest) GetFilePathQuotas() []*SetPolarFsFileQuotaRequestFilePathQuotas {
	return s.FilePathQuotas
}

func (s *SetPolarFsFileQuotaRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *SetPolarFsFileQuotaRequest) SetDBClusterId(v string) *SetPolarFsFileQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *SetPolarFsFileQuotaRequest) SetFilePathQuotas(v []*SetPolarFsFileQuotaRequestFilePathQuotas) *SetPolarFsFileQuotaRequest {
	s.FilePathQuotas = v
	return s
}

func (s *SetPolarFsFileQuotaRequest) SetPolarFsInstanceId(v string) *SetPolarFsFileQuotaRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *SetPolarFsFileQuotaRequest) Validate() error {
	if s.FilePathQuotas != nil {
		for _, item := range s.FilePathQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetPolarFsFileQuotaRequestFilePathQuotas struct {
	// example:
	//
	// /a/project
	FilePathId *string `json:"FilePathId,omitempty" xml:"FilePathId,omitempty"`
	// example:
	//
	// 1
	MaxDepth *int32 `json:"MaxDepth,omitempty" xml:"MaxDepth,omitempty"`
	// example:
	//
	// 1,2
	QuotaIds *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	// example:
	//
	// missing
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s SetPolarFsFileQuotaRequestFilePathQuotas) String() string {
	return dara.Prettify(s)
}

func (s SetPolarFsFileQuotaRequestFilePathQuotas) GoString() string {
	return s.String()
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) GetFilePathId() *string {
	return s.FilePathId
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) GetMaxDepth() *int32 {
	return s.MaxDepth
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) GetQuotaIds() *string {
	return s.QuotaIds
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) GetStrategy() *string {
	return s.Strategy
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) SetFilePathId(v string) *SetPolarFsFileQuotaRequestFilePathQuotas {
	s.FilePathId = &v
	return s
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) SetMaxDepth(v int32) *SetPolarFsFileQuotaRequestFilePathQuotas {
	s.MaxDepth = &v
	return s
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) SetQuotaIds(v string) *SetPolarFsFileQuotaRequestFilePathQuotas {
	s.QuotaIds = &v
	return s
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) SetStrategy(v string) *SetPolarFsFileQuotaRequestFilePathQuotas {
	s.Strategy = &v
	return s
}

func (s *SetPolarFsFileQuotaRequestFilePathQuotas) Validate() error {
	return dara.Validate(s)
}
