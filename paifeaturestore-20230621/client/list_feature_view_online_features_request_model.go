// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewOnlineFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJoinIds(v []*string) *ListFeatureViewOnlineFeaturesRequest
	GetJoinIds() []*string
}

type ListFeatureViewOnlineFeaturesRequest struct {
	// This parameter is required.
	JoinIds []*string `json:"JoinIds,omitempty" xml:"JoinIds,omitempty" type:"Repeated"`
}

func (s ListFeatureViewOnlineFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesRequest) GetJoinIds() []*string {
	return s.JoinIds
}

func (s *ListFeatureViewOnlineFeaturesRequest) SetJoinIds(v []*string) *ListFeatureViewOnlineFeaturesRequest {
	s.JoinIds = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
