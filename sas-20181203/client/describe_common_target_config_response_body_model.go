// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCommonTargetConfigResponseBody
	GetRequestId() *string
	SetTargetList(v []*DescribeCommonTargetConfigResponseBodyTargetList) *DescribeCommonTargetConfigResponseBody
	GetTargetList() []*DescribeCommonTargetConfigResponseBodyTargetList
}

type DescribeCommonTargetConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the configuration items.
	TargetList []*DescribeCommonTargetConfigResponseBodyTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
}

func (s DescribeCommonTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommonTargetConfigResponseBody) GetTargetList() []*DescribeCommonTargetConfigResponseBodyTargetList {
	return s.TargetList
}

func (s *DescribeCommonTargetConfigResponseBody) SetRequestId(v string) *DescribeCommonTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommonTargetConfigResponseBody) SetTargetList(v []*DescribeCommonTargetConfigResponseBodyTargetList) *DescribeCommonTargetConfigResponseBody {
	s.TargetList = v
	return s
}

func (s *DescribeCommonTargetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonTargetConfigResponseBodyTargetList struct {
	// The mode in which the configuration takes effect. Valid values:
	//
	// 	- **add**: In this mode, the configuration takes effect on the assets.
	//
	// 	- **del**: In this mode, the configuration does not take effect on the assets.
	//
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The ID of the asset on which the configuration takes effect.
	//
	// >
	//
	// 	- When you set the **TargetType*	- parameter to **uuid**, the value of this parameter indicates the UUID of an asset.
	//
	// 	- When you set the **TargetType*	- parameter to **Cluster**, the value of this parameter indicates the ID of a cluster.
	//
	// 	- When you set the **TargetType*	- parameter to **image_repo**, the value of this parameter indicates the ID of an image repository.
	//
	// example:
	//
	// c23551de6149343e8a54e69fbefe6****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The dimension from on which the feature was configured. Valid values:
	//
	// 	- **uuid**: the UUID of the asset
	//
	// 	- **Cluster**: the ID of the cluster
	//
	// 	- **image_repo**: the ID of the image repository
	//
	// example:
	//
	// image_repo
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeCommonTargetConfigResponseBodyTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetConfigResponseBodyTargetList) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) GetFlag() *string {
	return s.Flag
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) GetTarget() *string {
	return s.Target
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) SetFlag(v string) *DescribeCommonTargetConfigResponseBodyTargetList {
	s.Flag = &v
	return s
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) SetTarget(v string) *DescribeCommonTargetConfigResponseBodyTargetList {
	s.Target = &v
	return s
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) SetTargetType(v string) *DescribeCommonTargetConfigResponseBodyTargetList {
	s.TargetType = &v
	return s
}

func (s *DescribeCommonTargetConfigResponseBodyTargetList) Validate() error {
	return dara.Validate(s)
}
