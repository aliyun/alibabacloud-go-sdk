// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewOnlineFeaturesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ListFeatureViewOnlineFeaturesShrinkRequest
	GetConfig() *string
	SetJoinIdsShrink(v string) *ListFeatureViewOnlineFeaturesShrinkRequest
	GetJoinIdsShrink() *string
}

type ListFeatureViewOnlineFeaturesShrinkRequest struct {
	// example:
	//
	// {"dlrm_hstu": true}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	JoinIdsShrink *string `json:"JoinIds,omitempty" xml:"JoinIds,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) GetJoinIdsShrink() *string {
	return s.JoinIdsShrink
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) SetConfig(v string) *ListFeatureViewOnlineFeaturesShrinkRequest {
	s.Config = &v
	return s
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) SetJoinIdsShrink(v string) *ListFeatureViewOnlineFeaturesShrinkRequest {
	s.JoinIdsShrink = &v
	return s
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
