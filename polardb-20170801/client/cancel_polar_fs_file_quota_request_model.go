// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPolarFsFileQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelPolarFsFileQuotaRequest
	GetDBClusterId() *string
	SetFilePathIds(v string) *CancelPolarFsFileQuotaRequest
	GetFilePathIds() *string
	SetPolarFsInstanceId(v string) *CancelPolarFsFileQuotaRequest
	GetPolarFsInstanceId() *string
}

type CancelPolarFsFileQuotaRequest struct {
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /path1,/path2
	FilePathIds *string `json:"FilePathIds,omitempty" xml:"FilePathIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s CancelPolarFsFileQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPolarFsFileQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelPolarFsFileQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelPolarFsFileQuotaRequest) GetFilePathIds() *string {
	return s.FilePathIds
}

func (s *CancelPolarFsFileQuotaRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CancelPolarFsFileQuotaRequest) SetDBClusterId(v string) *CancelPolarFsFileQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelPolarFsFileQuotaRequest) SetFilePathIds(v string) *CancelPolarFsFileQuotaRequest {
	s.FilePathIds = &v
	return s
}

func (s *CancelPolarFsFileQuotaRequest) SetPolarFsInstanceId(v string) *CancelPolarFsFileQuotaRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CancelPolarFsFileQuotaRequest) Validate() error {
	return dara.Validate(s)
}
