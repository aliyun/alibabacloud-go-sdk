// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacUserCertStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*UpdateNacUserCertStatusRequestIdList) *UpdateNacUserCertStatusRequest
	GetIdList() []*UpdateNacUserCertStatusRequestIdList
	SetStatus(v string) *UpdateNacUserCertStatusRequest
	GetStatus() *string
}

type UpdateNacUserCertStatusRequest struct {
	IdList []*UpdateNacUserCertStatusRequestIdList `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateNacUserCertStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacUserCertStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusRequest) GetIdList() []*UpdateNacUserCertStatusRequestIdList {
	return s.IdList
}

func (s *UpdateNacUserCertStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateNacUserCertStatusRequest) SetIdList(v []*UpdateNacUserCertStatusRequestIdList) *UpdateNacUserCertStatusRequest {
	s.IdList = v
	return s
}

func (s *UpdateNacUserCertStatusRequest) SetStatus(v string) *UpdateNacUserCertStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateNacUserCertStatusRequest) Validate() error {
	if s.IdList != nil {
		for _, item := range s.IdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNacUserCertStatusRequestIdList struct {
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateNacUserCertStatusRequestIdList) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacUserCertStatusRequestIdList) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusRequestIdList) GetDevTag() *string {
	return s.DevTag
}

func (s *UpdateNacUserCertStatusRequestIdList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateNacUserCertStatusRequestIdList) SetDevTag(v string) *UpdateNacUserCertStatusRequestIdList {
	s.DevTag = &v
	return s
}

func (s *UpdateNacUserCertStatusRequestIdList) SetUserId(v string) *UpdateNacUserCertStatusRequestIdList {
	s.UserId = &v
	return s
}

func (s *UpdateNacUserCertStatusRequestIdList) Validate() error {
	return dara.Validate(s)
}
