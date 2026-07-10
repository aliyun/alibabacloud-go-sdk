// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertWhiteListSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNo(v string) *InsertWhiteListSettingRequest
	GetCertNo() *string
	SetCertifyId(v string) *InsertWhiteListSettingRequest
	GetCertifyId() *string
	SetRemark(v string) *InsertWhiteListSettingRequest
	GetRemark() *string
	SetSceneId(v int64) *InsertWhiteListSettingRequest
	GetSceneId() *int64
	SetServiceCode(v string) *InsertWhiteListSettingRequest
	GetServiceCode() *string
	SetValidDay(v int32) *InsertWhiteListSettingRequest
	GetValidDay() *int32
}

type InsertWhiteListSettingRequest struct {
	// The certificate number.
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The unique identifier of the ID Verification request.
	//
	// example:
	//
	// shsf57a4e0d9981c3bd66dc754f3d3cd
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The remarks. The maximum length is 32 characters.
	//
	// example:
	//
	// xxxxxx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The verification scenario ID. This ID is automatically generated after you create a verification scenario in the console. For information about how to create a verification scenario, see Add a verification scenario.
	//
	// example:
	//
	// 100000xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The ServiceCode of the ID Verification cloud service. Set this parameter to **antcloudauth**.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The number of days the whitelist entry is valid. Valid values: 3, 7, and 30.
	//
	// example:
	//
	// 30
	ValidDay *int32 `json:"ValidDay,omitempty" xml:"ValidDay,omitempty"`
}

func (s InsertWhiteListSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *InsertWhiteListSettingRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InsertWhiteListSettingRequest) GetRemark() *string {
	return s.Remark
}

func (s *InsertWhiteListSettingRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *InsertWhiteListSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *InsertWhiteListSettingRequest) GetValidDay() *int32 {
	return s.ValidDay
}

func (s *InsertWhiteListSettingRequest) SetCertNo(v string) *InsertWhiteListSettingRequest {
	s.CertNo = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetCertifyId(v string) *InsertWhiteListSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetRemark(v string) *InsertWhiteListSettingRequest {
	s.Remark = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetSceneId(v int64) *InsertWhiteListSettingRequest {
	s.SceneId = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetServiceCode(v string) *InsertWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *InsertWhiteListSettingRequest) SetValidDay(v int32) *InsertWhiteListSettingRequest {
	s.ValidDay = &v
	return s
}

func (s *InsertWhiteListSettingRequest) Validate() error {
	return dara.Validate(s)
}
