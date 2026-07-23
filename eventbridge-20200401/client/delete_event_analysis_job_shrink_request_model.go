// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventAnalysisJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceResourceShrink(v string) *DeleteEventAnalysisJobShrinkRequest
	GetSourceResourceShrink() *string
}

type DeleteEventAnalysisJobShrinkRequest struct {
	// The identifier of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Kafka":{"RegionId":"cn-hangzhou","InstanceId":"alikafka_post-cn-xxx","Topic":"my_topic"}}
	SourceResourceShrink *string `json:"SourceResource,omitempty" xml:"SourceResource,omitempty"`
}

func (s DeleteEventAnalysisJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobShrinkRequest) GetSourceResourceShrink() *string {
	return s.SourceResourceShrink
}

func (s *DeleteEventAnalysisJobShrinkRequest) SetSourceResourceShrink(v string) *DeleteEventAnalysisJobShrinkRequest {
	s.SourceResourceShrink = &v
	return s
}

func (s *DeleteEventAnalysisJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
