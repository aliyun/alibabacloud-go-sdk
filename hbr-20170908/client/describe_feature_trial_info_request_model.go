// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureTrialInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v string) *DescribeFeatureTrialInfoRequest
	GetFeatureType() *string
}

type DescribeFeatureTrialInfoRequest struct {
	// The feature type. Currently, only the free trial information of Tablestore backup can be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// OTS_BACKUP
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
}

func (s DescribeFeatureTrialInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTrialInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTrialInfoRequest) GetFeatureType() *string {
	return s.FeatureType
}

func (s *DescribeFeatureTrialInfoRequest) SetFeatureType(v string) *DescribeFeatureTrialInfoRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeFeatureTrialInfoRequest) Validate() error {
	return dara.Validate(s)
}
