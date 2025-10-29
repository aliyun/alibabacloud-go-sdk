// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolarFsFileQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *SetPolarFsFileQuotaResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *SetPolarFsFileQuotaResponseBody
	GetRequestId() *string
}

type SetPolarFsFileQuotaResponseBody struct {
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolarFsFileQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolarFsFileQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolarFsFileQuotaResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *SetPolarFsFileQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolarFsFileQuotaResponseBody) SetPolarFsInstanceId(v string) *SetPolarFsFileQuotaResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *SetPolarFsFileQuotaResponseBody) SetRequestId(v string) *SetPolarFsFileQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolarFsFileQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
