// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDistributeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageDistributeInfoResponseBody
	GetCode() *string
	SetData(v *GetImageDistributeInfoResponseBodyData) *GetImageDistributeInfoResponseBody
	GetData() *GetImageDistributeInfoResponseBodyData
	SetMessage(v string) *GetImageDistributeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageDistributeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageDistributeInfoResponseBody
	GetSuccess() *bool
}

type GetImageDistributeInfoResponseBody struct {
	// The return code of the API. A value of 200 indicates success. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image information returned by the query.
	Data *GetImageDistributeInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The processing information returned by the API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values: true: Successful. false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageDistributeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageDistributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDistributeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageDistributeInfoResponseBody) GetData() *GetImageDistributeInfoResponseBodyData {
	return s.Data
}

func (s *GetImageDistributeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageDistributeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageDistributeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageDistributeInfoResponseBody) SetCode(v string) *GetImageDistributeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDistributeInfoResponseBody) SetData(v *GetImageDistributeInfoResponseBodyData) *GetImageDistributeInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetImageDistributeInfoResponseBody) SetMessage(v string) *GetImageDistributeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDistributeInfoResponseBody) SetRequestId(v string) *GetImageDistributeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBody) SetSuccess(v bool) *GetImageDistributeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageDistributeInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDistributeInfoResponseBodyData struct {
	// The list of region distribution information for the image version.
	ImageRegionDistributeList []*GetImageDistributeInfoResponseBodyDataImageRegionDistributeList `json:"ImageRegionDistributeList,omitempty" xml:"ImageRegionDistributeList,omitempty" type:"Repeated"`
	// The list of region IDs supported by the image.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The source of the image. For example, center indicates the central image service.
	//
	// example:
	//
	// center
	ImageSource *string `json:"ImageSource,omitempty" xml:"ImageSource,omitempty"`
}

func (s GetImageDistributeInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageDistributeInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageDistributeInfoResponseBodyData) GetImageRegionDistributeList() []*GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	return s.ImageRegionDistributeList
}

func (s *GetImageDistributeInfoResponseBodyData) GetImageRegionList() []*string {
	return s.ImageRegionList
}

func (s *GetImageDistributeInfoResponseBodyData) GetImageSource() *string {
	return s.ImageSource
}

func (s *GetImageDistributeInfoResponseBodyData) SetImageRegionDistributeList(v []*GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) *GetImageDistributeInfoResponseBodyData {
	s.ImageRegionDistributeList = v
	return s
}

func (s *GetImageDistributeInfoResponseBodyData) SetImageRegionList(v []*string) *GetImageDistributeInfoResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *GetImageDistributeInfoResponseBodyData) SetImageSource(v string) *GetImageDistributeInfoResponseBodyData {
	s.ImageSource = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyData) Validate() error {
	if s.ImageRegionDistributeList != nil {
		for _, item := range s.ImageRegionDistributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageDistributeInfoResponseBodyDataImageRegionDistributeList struct {
	// The ID of the Alibaba Cloud account to which the image belongs. The account ID for a system image is 0.
	//
	// example:
	//
	// 123456789
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the ECS image corresponding to the region.
	//
	// example:
	//
	// m-bp1example****
	EcsImageId *string `json:"EcsImageId,omitempty" xml:"EcsImageId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The processing information reported by the region distribution task.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The distribution progress of the image in the region, expressed as a percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region to which the image is distributed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region distribution status.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logical ID of the image on the unit side, which can be different from the central image ID.
	//
	// example:
	//
	// m-1bdedef****
	UnitImageId *string `json:"UnitImageId,omitempty" xml:"UnitImageId,omitempty"`
	// The image version ID.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) String() string {
	return dara.Prettify(s)
}

func (s GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GoString() string {
	return s.String()
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetEcsImageId() *string {
	return s.EcsImageId
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetMessage() *string {
	return s.Message
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetProgress() *string {
	return s.Progress
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetStatus() *string {
	return s.Status
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetUnitImageId() *string {
	return s.UnitImageId
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetAliUid(v int64) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.AliUid = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetEcsImageId(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.EcsImageId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetGmtCreate(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.GmtCreate = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetGmtModified(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.GmtModified = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetImageId(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.ImageId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetMessage(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.Message = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetProgress(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.Progress = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetRegionId(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.RegionId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetStatus(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.Status = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetUnitImageId(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.UnitImageId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) SetVersionId(v string) *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList {
	s.VersionId = &v
	return s
}

func (s *GetImageDistributeInfoResponseBodyDataImageRegionDistributeList) Validate() error {
	return dara.Validate(s)
}
