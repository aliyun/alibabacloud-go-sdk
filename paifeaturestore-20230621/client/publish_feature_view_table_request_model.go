// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFeatureViewTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *PublishFeatureViewTableRequest
	GetConfig() *string
	SetEventTime(v string) *PublishFeatureViewTableRequest
	GetEventTime() *string
	SetMode(v string) *PublishFeatureViewTableRequest
	GetMode() *string
	SetOfflineToOnline(v bool) *PublishFeatureViewTableRequest
	GetOfflineToOnline() *bool
	SetPartitions(v map[string]map[string]interface{}) *PublishFeatureViewTableRequest
	GetPartitions() map[string]map[string]interface{}
}

type PublishFeatureViewTableRequest struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Overwrite
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OfflineToOnline *bool                             `json:"OfflineToOnline,omitempty" xml:"OfflineToOnline,omitempty"`
	Partitions      map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s PublishFeatureViewTableRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishFeatureViewTableRequest) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableRequest) GetConfig() *string {
	return s.Config
}

func (s *PublishFeatureViewTableRequest) GetEventTime() *string {
	return s.EventTime
}

func (s *PublishFeatureViewTableRequest) GetMode() *string {
	return s.Mode
}

func (s *PublishFeatureViewTableRequest) GetOfflineToOnline() *bool {
	return s.OfflineToOnline
}

func (s *PublishFeatureViewTableRequest) GetPartitions() map[string]map[string]interface{} {
	return s.Partitions
}

func (s *PublishFeatureViewTableRequest) SetConfig(v string) *PublishFeatureViewTableRequest {
	s.Config = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetEventTime(v string) *PublishFeatureViewTableRequest {
	s.EventTime = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetMode(v string) *PublishFeatureViewTableRequest {
	s.Mode = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetOfflineToOnline(v bool) *PublishFeatureViewTableRequest {
	s.OfflineToOnline = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetPartitions(v map[string]map[string]interface{}) *PublishFeatureViewTableRequest {
	s.Partitions = v
	return s
}

func (s *PublishFeatureViewTableRequest) Validate() error {
	return dara.Validate(s)
}
