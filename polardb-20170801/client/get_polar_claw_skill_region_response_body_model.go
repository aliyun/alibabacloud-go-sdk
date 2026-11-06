// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawSkillRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOSSBucket(v string) *GetPolarClawSkillRegionResponseBody
	GetOSSBucket() *string
	SetRegionId(v string) *GetPolarClawSkillRegionResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetPolarClawSkillRegionResponseBody
	GetRequestId() *string
}

type GetPolarClawSkillRegionResponseBody struct {
	// The OSS bucket that stores images.
	//
	// example:
	//
	// oss-bucket-test
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolarClawSkillRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawSkillRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarClawSkillRegionResponseBody) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *GetPolarClawSkillRegionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolarClawSkillRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolarClawSkillRegionResponseBody) SetOSSBucket(v string) *GetPolarClawSkillRegionResponseBody {
	s.OSSBucket = &v
	return s
}

func (s *GetPolarClawSkillRegionResponseBody) SetRegionId(v string) *GetPolarClawSkillRegionResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetPolarClawSkillRegionResponseBody) SetRequestId(v string) *GetPolarClawSkillRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolarClawSkillRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
