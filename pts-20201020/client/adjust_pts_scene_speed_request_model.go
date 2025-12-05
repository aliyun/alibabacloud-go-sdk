// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustPtsSceneSpeedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiSpeedList(v []*AdjustPtsSceneSpeedRequestApiSpeedList) *AdjustPtsSceneSpeedRequest
	GetApiSpeedList() []*AdjustPtsSceneSpeedRequestApiSpeedList
	SetSceneId(v string) *AdjustPtsSceneSpeedRequest
	GetSceneId() *string
}

type AdjustPtsSceneSpeedRequest struct {
	// The stress testing speed in the PTS scenario.
	ApiSpeedList []*AdjustPtsSceneSpeedRequestApiSpeedList `json:"ApiSpeedList,omitempty" xml:"ApiSpeedList,omitempty" type:"Repeated"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYXXX12H
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AdjustPtsSceneSpeedRequest) String() string {
	return dara.Prettify(s)
}

func (s AdjustPtsSceneSpeedRequest) GoString() string {
	return s.String()
}

func (s *AdjustPtsSceneSpeedRequest) GetApiSpeedList() []*AdjustPtsSceneSpeedRequestApiSpeedList {
	return s.ApiSpeedList
}

func (s *AdjustPtsSceneSpeedRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AdjustPtsSceneSpeedRequest) SetApiSpeedList(v []*AdjustPtsSceneSpeedRequestApiSpeedList) *AdjustPtsSceneSpeedRequest {
	s.ApiSpeedList = v
	return s
}

func (s *AdjustPtsSceneSpeedRequest) SetSceneId(v string) *AdjustPtsSceneSpeedRequest {
	s.SceneId = &v
	return s
}

func (s *AdjustPtsSceneSpeedRequest) Validate() error {
	if s.ApiSpeedList != nil {
		for _, item := range s.ApiSpeedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AdjustPtsSceneSpeedRequestApiSpeedList struct {
	// The API ID. You can find the information of the API corresponding to the ID in the Relation response parameter of the GetPtsSceneRunningData operation based on the ID.
	//
	// example:
	//
	// DYXXX12H
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The new stress. In concurrency mode, the new stress is the concurrency. In RPS mode, the new stress is the RPS.
	//
	// example:
	//
	// 30
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s AdjustPtsSceneSpeedRequestApiSpeedList) String() string {
	return dara.Prettify(s)
}

func (s AdjustPtsSceneSpeedRequestApiSpeedList) GoString() string {
	return s.String()
}

func (s *AdjustPtsSceneSpeedRequestApiSpeedList) GetApiId() *string {
	return s.ApiId
}

func (s *AdjustPtsSceneSpeedRequestApiSpeedList) GetSpeed() *int64 {
	return s.Speed
}

func (s *AdjustPtsSceneSpeedRequestApiSpeedList) SetApiId(v string) *AdjustPtsSceneSpeedRequestApiSpeedList {
	s.ApiId = &v
	return s
}

func (s *AdjustPtsSceneSpeedRequestApiSpeedList) SetSpeed(v int64) *AdjustPtsSceneSpeedRequestApiSpeedList {
	s.Speed = &v
	return s
}

func (s *AdjustPtsSceneSpeedRequestApiSpeedList) Validate() error {
	return dara.Validate(s)
}
