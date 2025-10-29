// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *DeletePolarFsQuotaResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *DeletePolarFsQuotaResponseBody
	GetRequestId() *string
}

type DeletePolarFsQuotaResponseBody struct {
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarFsQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarFsQuotaResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarFsQuotaResponseBody) SetPolarFsInstanceId(v string) *DeletePolarFsQuotaResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsQuotaResponseBody) SetRequestId(v string) *DeletePolarFsQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarFsQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
