// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskWarningLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDiskWarningLineRequest
	GetClusterId() *string
}

type DescribeDiskWarningLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1bl7iqzkahmyxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDiskWarningLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDiskWarningLineRequest) SetClusterId(v string) *DescribeDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDiskWarningLineRequest) Validate() error {
	return dara.Validate(s)
}
