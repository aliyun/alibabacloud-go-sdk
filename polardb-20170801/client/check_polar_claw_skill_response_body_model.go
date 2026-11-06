// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v []*string) *CheckPolarClawSkillResponseBody
	GetMessage() []*string
	SetOSSBucket(v string) *CheckPolarClawSkillResponseBody
	GetOSSBucket() *string
	SetOSSRegionId(v string) *CheckPolarClawSkillResponseBody
	GetOSSRegionId() *string
	SetRequestId(v string) *CheckPolarClawSkillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckPolarClawSkillResponseBody
	GetSuccess() *bool
}

type CheckPolarClawSkillResponseBody struct {
	// The returned message.
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
	// The name of the OSS bucket where the skill files are stored.
	//
	// example:
	//
	// oss-bucket-test
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The region ID of the OSS bucket where the skill files are stored.
	//
	// example:
	//
	// cn-hangzhou
	OSSRegionId *string `json:"OSSRegionId,omitempty" xml:"OSSRegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckPolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CheckPolarClawSkillResponseBody) GetMessage() []*string {
	return s.Message
}

func (s *CheckPolarClawSkillResponseBody) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *CheckPolarClawSkillResponseBody) GetOSSRegionId() *string {
	return s.OSSRegionId
}

func (s *CheckPolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckPolarClawSkillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckPolarClawSkillResponseBody) SetMessage(v []*string) *CheckPolarClawSkillResponseBody {
	s.Message = v
	return s
}

func (s *CheckPolarClawSkillResponseBody) SetOSSBucket(v string) *CheckPolarClawSkillResponseBody {
	s.OSSBucket = &v
	return s
}

func (s *CheckPolarClawSkillResponseBody) SetOSSRegionId(v string) *CheckPolarClawSkillResponseBody {
	s.OSSRegionId = &v
	return s
}

func (s *CheckPolarClawSkillResponseBody) SetRequestId(v string) *CheckPolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckPolarClawSkillResponseBody) SetSuccess(v bool) *CheckPolarClawSkillResponseBody {
	s.Success = &v
	return s
}

func (s *CheckPolarClawSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
