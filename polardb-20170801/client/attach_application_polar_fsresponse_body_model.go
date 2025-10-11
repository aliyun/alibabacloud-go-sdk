// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplicationPolarFSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AttachApplicationPolarFSResponseBody
	GetApplicationId() *string
	SetPolarFSInstanceId(v string) *AttachApplicationPolarFSResponseBody
	GetPolarFSInstanceId() *string
	SetRequestId(v string) *AttachApplicationPolarFSResponseBody
	GetRequestId() *string
}

type AttachApplicationPolarFSResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// PolarFS ID
	//
	// example:
	//
	// pcs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachApplicationPolarFSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachApplicationPolarFSResponseBody) GoString() string {
	return s.String()
}

func (s *AttachApplicationPolarFSResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AttachApplicationPolarFSResponseBody) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *AttachApplicationPolarFSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachApplicationPolarFSResponseBody) SetApplicationId(v string) *AttachApplicationPolarFSResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *AttachApplicationPolarFSResponseBody) SetPolarFSInstanceId(v string) *AttachApplicationPolarFSResponseBody {
	s.PolarFSInstanceId = &v
	return s
}

func (s *AttachApplicationPolarFSResponseBody) SetRequestId(v string) *AttachApplicationPolarFSResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachApplicationPolarFSResponseBody) Validate() error {
	return dara.Validate(s)
}
