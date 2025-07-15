// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *GetEditingJobInfoRequest
	GetCasterId() *string
	SetOwnerId(v int64) *GetEditingJobInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEditingJobInfoRequest
	GetRegionId() *string
	SetShowId(v string) *GetEditingJobInfoRequest
	GetShowId() *string
}

type GetEditingJobInfoRequest struct {
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the episode for which you want to query editing tasks.
	//
	// >  You can obtain the ID from the response parameter ShowId of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/370861.html) operation.
	//
	// example:
	//
	// 72200b81-b761-4c10-842a-a0726d97****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
}

func (s GetEditingJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEditingJobInfoRequest) GoString() string {
	return s.String()
}

func (s *GetEditingJobInfoRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *GetEditingJobInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEditingJobInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEditingJobInfoRequest) GetShowId() *string {
	return s.ShowId
}

func (s *GetEditingJobInfoRequest) SetCasterId(v string) *GetEditingJobInfoRequest {
	s.CasterId = &v
	return s
}

func (s *GetEditingJobInfoRequest) SetOwnerId(v int64) *GetEditingJobInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEditingJobInfoRequest) SetRegionId(v string) *GetEditingJobInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetEditingJobInfoRequest) SetShowId(v string) *GetEditingJobInfoRequest {
	s.ShowId = &v
	return s
}

func (s *GetEditingJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
