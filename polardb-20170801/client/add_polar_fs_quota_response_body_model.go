// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarFsQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *AddPolarFsQuotaResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *AddPolarFsQuotaResponseBody
	GetRequestId() *string
}

type AddPolarFsQuotaResponseBody struct {
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F2007D3-7E74-4ECB-89A8-BF130D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPolarFsQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *AddPolarFsQuotaResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *AddPolarFsQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPolarFsQuotaResponseBody) SetPolarFsInstanceId(v string) *AddPolarFsQuotaResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *AddPolarFsQuotaResponseBody) SetRequestId(v string) *AddPolarFsQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPolarFsQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
