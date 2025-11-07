// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkReachableAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPathId(v string) *CreateNetworkReachableAnalysisRequest
	GetNetworkPathId() *string
	SetRegionId(v string) *CreateNetworkReachableAnalysisRequest
	GetRegionId() *string
	SetTag(v []*CreateNetworkReachableAnalysisRequestTag) *CreateNetworkReachableAnalysisRequest
	GetTag() []*CreateNetworkReachableAnalysisRequestTag
}

type CreateNetworkReachableAnalysisRequest struct {
	// The ID of the network path. You can call the [CreateNetworkPath](https://help.aliyun.com/document_detail/2366522.html) operation to obtain the ID of the network path.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-b2f618ceb2c84057****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The ID of the region for which you want to create a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateNetworkReachableAnalysisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNetworkReachableAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisRequest) GetNetworkPathId() *string {
	return s.NetworkPathId
}

func (s *CreateNetworkReachableAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkReachableAnalysisRequest) GetTag() []*CreateNetworkReachableAnalysisRequestTag {
	return s.Tag
}

func (s *CreateNetworkReachableAnalysisRequest) SetNetworkPathId(v string) *CreateNetworkReachableAnalysisRequest {
	s.NetworkPathId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequest) SetRegionId(v string) *CreateNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequest) SetTag(v []*CreateNetworkReachableAnalysisRequestTag) *CreateNetworkReachableAnalysisRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkReachableAnalysisRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkReachableAnalysisRequestTag struct {
	// The key of the tag to add to the resource. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// You can add up to 20 tags in each call.
	//
	// example:
	//
	// Team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`. The tag value can be an empty string.
	//
	// You can add up to 20 tag values in each call.
	//
	// example:
	//
	// ops
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkReachableAnalysisRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkReachableAnalysisRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkReachableAnalysisRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkReachableAnalysisRequestTag) SetKey(v string) *CreateNetworkReachableAnalysisRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequestTag) SetValue(v string) *CreateNetworkReachableAnalysisRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequestTag) Validate() error {
	return dara.Validate(s)
}
