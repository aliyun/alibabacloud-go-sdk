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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter in the response.
	//
	// - If you created the production studio in the ApsaraVideo Live console, navigate to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > - The name of the production studio in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// > - CasterId must be a production studio with NormType=6 (playlist mode). Using a production studio with other NormType values (such as 1 or 3) returns InvalidShowList.NotFound. You can filter by NormType=6 in the DescribeCasters response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the show to query.
	//
	// >You can obtain the ShowId value from the response of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/370861.html) operation.
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
