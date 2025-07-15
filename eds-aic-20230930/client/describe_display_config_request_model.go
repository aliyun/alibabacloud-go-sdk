// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisplayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DescribeDisplayConfigRequest
	GetAndroidInstanceIds() []*string
}

type DescribeDisplayConfigRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDisplayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisplayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DescribeDisplayConfigRequest) SetAndroidInstanceIds(v []*string) *DescribeDisplayConfigRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeDisplayConfigRequest) Validate() error {
	return dara.Validate(s)
}
