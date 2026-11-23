// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyActivityAdmissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *VerifyActivityAdmissionRequest
	GetActivityId() *int64
	SetIdCardNo(v string) *VerifyActivityAdmissionRequest
	GetIdCardNo() *string
}

type VerifyActivityAdmissionRequest struct {
	// This parameter is required.
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	IdCardNo *string `json:"IdCardNo,omitempty" xml:"IdCardNo,omitempty"`
}

func (s VerifyActivityAdmissionRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyActivityAdmissionRequest) GoString() string {
	return s.String()
}

func (s *VerifyActivityAdmissionRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *VerifyActivityAdmissionRequest) GetIdCardNo() *string {
	return s.IdCardNo
}

func (s *VerifyActivityAdmissionRequest) SetActivityId(v int64) *VerifyActivityAdmissionRequest {
	s.ActivityId = &v
	return s
}

func (s *VerifyActivityAdmissionRequest) SetIdCardNo(v string) *VerifyActivityAdmissionRequest {
	s.IdCardNo = &v
	return s
}

func (s *VerifyActivityAdmissionRequest) Validate() error {
	return dara.Validate(s)
}
