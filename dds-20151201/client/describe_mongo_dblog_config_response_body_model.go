// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMongoDBLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAudit(v bool) *DescribeMongoDBLogConfigResponseBody
	GetEnableAudit() *bool
	SetIsEtlMetaExist(v int32) *DescribeMongoDBLogConfigResponseBody
	GetIsEtlMetaExist() *int32
	SetIsUserProjectLogstoreExist(v int32) *DescribeMongoDBLogConfigResponseBody
	GetIsUserProjectLogstoreExist() *int32
	SetPreserveStorageForStandard(v int64) *DescribeMongoDBLogConfigResponseBody
	GetPreserveStorageForStandard() *int64
	SetPreserveStorageForTrail(v int64) *DescribeMongoDBLogConfigResponseBody
	GetPreserveStorageForTrail() *int64
	SetRequestId(v string) *DescribeMongoDBLogConfigResponseBody
	GetRequestId() *string
	SetServiceType(v string) *DescribeMongoDBLogConfigResponseBody
	GetServiceType() *string
	SetTtlForStandard(v int64) *DescribeMongoDBLogConfigResponseBody
	GetTtlForStandard() *int64
	SetTtlForTrail(v int64) *DescribeMongoDBLogConfigResponseBody
	GetTtlForTrail() *int64
	SetUsedStorageForStandard(v int64) *DescribeMongoDBLogConfigResponseBody
	GetUsedStorageForStandard() *int64
	SetUsedStorageForTrail(v int64) *DescribeMongoDBLogConfigResponseBody
	GetUsedStorageForTrail() *int64
	SetUserProjectName(v string) *DescribeMongoDBLogConfigResponseBody
	GetUserProjectName() *string
}

type DescribeMongoDBLogConfigResponseBody struct {
	// Indicates whether to enable the audit log feature.
	//
	// 	- **true**: The audit log feature is enabled.
	//
	// 	- **false**: The audit log feature is disabled.
	//
	// example:
	//
	// true
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// Indicates whether a rule to distribute logs to Logtail is created. For more information, see [Logtail overview](https://help.aliyun.com/document_detail/28979.html). Valid values:
	//
	// 	- **1**: A rule to distribute logs to Logtail is created.
	//
	// 	- **0*	- or **null**: A rule to distribute logs to Logtail is not created.
	//
	// example:
	//
	// 1
	IsEtlMetaExist *int32 `json:"IsEtlMetaExist,omitempty" xml:"IsEtlMetaExist,omitempty"`
	// Indicates whether a project exists in the current region. Valid values:
	//
	// 	- **1**: A logging project exists in the current region.
	//
	// 	- **0*	- or **null**: A logging project does not exist in the current region.
	//
	// example:
	//
	// 1
	IsUserProjectLogstoreExist *int32 `json:"IsUserProjectLogstoreExist,omitempty" xml:"IsUserProjectLogstoreExist,omitempty"`
	// The maximum storage capacity for the formal edition of the audit log feature. If the value is -1, no maximum storage capacity is set.
	//
	// example:
	//
	// -1
	PreserveStorageForStandard *int64 `json:"PreserveStorageForStandard,omitempty" xml:"PreserveStorageForStandard,omitempty"`
	// The maximum storage capacity for the free trial edition of the audit log feature. Unit: bytes. You can set the maximum storage capacity to 107,374,182,400 bytes.
	//
	// example:
	//
	// 107374182400
	PreserveStorageForTrail *int64 `json:"PreserveStorageForTrail,omitempty" xml:"PreserveStorageForTrail,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 664ECE26-658A-47C5-88F6-870B0132E8D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the audit log feature. Valid values:
	//
	// 	- **Trail**: the free trial edition
	//
	// 	- **Standard**: the official edition
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The retention period for the official edition of the audit log feature. Valid values: 1 to 365. Unit: day.
	//
	// example:
	//
	// 30
	TtlForStandard *int64 `json:"TtlForStandard,omitempty" xml:"TtlForStandard,omitempty"`
	// The retention period for the free trial edition of the audit log feature.
	//
	// example:
	//
	// 1
	TtlForTrail *int64 `json:"TtlForTrail,omitempty" xml:"TtlForTrail,omitempty"`
	// The used storage capacity for the formal edition of the audit log feature. Unit: bytes.
	//
	// example:
	//
	// 20163
	UsedStorageForStandard *int64 `json:"UsedStorageForStandard,omitempty" xml:"UsedStorageForStandard,omitempty"`
	// The used storage capacity for the free trial edition of the audit log feature. Unit: bytes.
	//
	// example:
	//
	// 12548178759
	UsedStorageForTrail *int64 `json:"UsedStorageForTrail,omitempty" xml:"UsedStorageForTrail,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// nosql-176498472570****-cn-hangzhou
	UserProjectName *string `json:"UserProjectName,omitempty" xml:"UserProjectName,omitempty"`
}

func (s DescribeMongoDBLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMongoDBLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigResponseBody) GetEnableAudit() *bool {
	return s.EnableAudit
}

func (s *DescribeMongoDBLogConfigResponseBody) GetIsEtlMetaExist() *int32 {
	return s.IsEtlMetaExist
}

func (s *DescribeMongoDBLogConfigResponseBody) GetIsUserProjectLogstoreExist() *int32 {
	return s.IsUserProjectLogstoreExist
}

func (s *DescribeMongoDBLogConfigResponseBody) GetPreserveStorageForStandard() *int64 {
	return s.PreserveStorageForStandard
}

func (s *DescribeMongoDBLogConfigResponseBody) GetPreserveStorageForTrail() *int64 {
	return s.PreserveStorageForTrail
}

func (s *DescribeMongoDBLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMongoDBLogConfigResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeMongoDBLogConfigResponseBody) GetTtlForStandard() *int64 {
	return s.TtlForStandard
}

func (s *DescribeMongoDBLogConfigResponseBody) GetTtlForTrail() *int64 {
	return s.TtlForTrail
}

func (s *DescribeMongoDBLogConfigResponseBody) GetUsedStorageForStandard() *int64 {
	return s.UsedStorageForStandard
}

func (s *DescribeMongoDBLogConfigResponseBody) GetUsedStorageForTrail() *int64 {
	return s.UsedStorageForTrail
}

func (s *DescribeMongoDBLogConfigResponseBody) GetUserProjectName() *string {
	return s.UserProjectName
}

func (s *DescribeMongoDBLogConfigResponseBody) SetEnableAudit(v bool) *DescribeMongoDBLogConfigResponseBody {
	s.EnableAudit = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetIsEtlMetaExist(v int32) *DescribeMongoDBLogConfigResponseBody {
	s.IsEtlMetaExist = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetIsUserProjectLogstoreExist(v int32) *DescribeMongoDBLogConfigResponseBody {
	s.IsUserProjectLogstoreExist = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetPreserveStorageForStandard(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.PreserveStorageForStandard = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetPreserveStorageForTrail(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.PreserveStorageForTrail = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetRequestId(v string) *DescribeMongoDBLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetServiceType(v string) *DescribeMongoDBLogConfigResponseBody {
	s.ServiceType = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetTtlForStandard(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.TtlForStandard = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetTtlForTrail(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.TtlForTrail = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetUsedStorageForStandard(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.UsedStorageForStandard = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetUsedStorageForTrail(v int64) *DescribeMongoDBLogConfigResponseBody {
	s.UsedStorageForTrail = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetUserProjectName(v string) *DescribeMongoDBLogConfigResponseBody {
	s.UserProjectName = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
