// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPolarFsFileQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilePathIds(v string) *CancelPolarFsFileQuotaResponseBody
	GetFilePathIds() *string
	SetPolarFsInstanceId(v string) *CancelPolarFsFileQuotaResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *CancelPolarFsFileQuotaResponseBody
	GetRequestId() *string
}

type CancelPolarFsFileQuotaResponseBody struct {
	// example:
	//
	// /path1,/path2
	FilePathIds *string `json:"FilePathIds,omitempty" xml:"FilePathIds,omitempty"`
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPolarFsFileQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPolarFsFileQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPolarFsFileQuotaResponseBody) GetFilePathIds() *string {
	return s.FilePathIds
}

func (s *CancelPolarFsFileQuotaResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CancelPolarFsFileQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPolarFsFileQuotaResponseBody) SetFilePathIds(v string) *CancelPolarFsFileQuotaResponseBody {
	s.FilePathIds = &v
	return s
}

func (s *CancelPolarFsFileQuotaResponseBody) SetPolarFsInstanceId(v string) *CancelPolarFsFileQuotaResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CancelPolarFsFileQuotaResponseBody) SetRequestId(v string) *CancelPolarFsFileQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPolarFsFileQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
