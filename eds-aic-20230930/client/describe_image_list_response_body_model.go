// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeImageListResponseBodyData) *DescribeImageListResponseBody
	GetData() []*DescribeImageListResponseBodyData
	SetNextToken(v string) *DescribeImageListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageListResponseBody
	GetTotalCount() *int32
}

type DescribeImageListResponseBody struct {
	// The images.
	Data []*DescribeImageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 620740FF-492F-5956-B1BA-361E966C0269
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponseBody) GetData() []*DescribeImageListResponseBodyData {
	return s.Data
}

func (s *DescribeImageListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageListResponseBody) SetData(v []*DescribeImageListResponseBodyData) *DescribeImageListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageListResponseBody) SetNextToken(v string) *DescribeImageListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImageListResponseBody) SetRequestId(v string) *DescribeImageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageListResponseBody) SetTotalCount(v int32) *DescribeImageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageListResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 117819727354****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the image was last modified.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtModified  *string                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageBizTags []*DescribeImageListResponseBodyDataImageBizTags `json:"ImageBizTags,omitempty" xml:"ImageBizTags,omitempty" type:"Repeated"`
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// IMAGE
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The region where the image is distributed. The key is the region and the value is the distribution information.
	ImageRegionDistributeMap map[string]*DataImageRegionDistributeMapValue `json:"ImageRegionDistributeMap,omitempty" xml:"ImageRegionDistributeMap,omitempty"`
	// The list of regions.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The type of the image.
	//
	// Valid values:
	//
	// 	- User: custom images.
	//
	// 	- System: system images.
	//
	// example:
	//
	// System
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The language of the image.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time when the image was published.
	//
	// example:
	//
	// 2024-07-25 10:06:45
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The rendering type.
	//
	// Valid values:
	//
	// 	- GPURemote
	//
	// 	- CPU
	//
	// 	- GPULocal
	//
	// example:
	//
	// CPU
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The state of the image.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The image is available.
	//
	// 	- DELETE: The image is deleted.
	//
	// 	- INIT: The image is being initialized.
	//
	// 	- CREATE_FAILED: The image failed to be created.
	//
	// 	- CREATING: The image is being created.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The OS type of the image.
	//
	// example:
	//
	// Android 12
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s DescribeImageListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeImageListResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageListResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeImageListResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeImageListResponseBodyData) GetImageBizTags() []*DescribeImageListResponseBodyDataImageBizTags {
	return s.ImageBizTags
}

func (s *DescribeImageListResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageListResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImageListResponseBodyData) GetImageRegionDistributeMap() map[string]*DataImageRegionDistributeMapValue {
	return s.ImageRegionDistributeMap
}

func (s *DescribeImageListResponseBodyData) GetImageRegionList() []*string {
	return s.ImageRegionList
}

func (s *DescribeImageListResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeImageListResponseBodyData) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeImageListResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *DescribeImageListResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeImageListResponseBodyData) GetRenderingType() *string {
	return s.RenderingType
}

func (s *DescribeImageListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageListResponseBodyData) GetSystemType() *string {
	return s.SystemType
}

func (s *DescribeImageListResponseBodyData) SetAliUid(v int64) *DescribeImageListResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetDescription(v string) *DescribeImageListResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetGmtCreate(v string) *DescribeImageListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetGmtModified(v string) *DescribeImageListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageBizTags(v []*DescribeImageListResponseBodyDataImageBizTags) *DescribeImageListResponseBodyData {
	s.ImageBizTags = v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageId(v string) *DescribeImageListResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageName(v string) *DescribeImageListResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageRegionDistributeMap(v map[string]*DataImageRegionDistributeMapValue) *DescribeImageListResponseBodyData {
	s.ImageRegionDistributeMap = v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageRegionList(v []*string) *DescribeImageListResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageType(v string) *DescribeImageListResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageVersion(v string) *DescribeImageListResponseBodyData {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetLanguage(v string) *DescribeImageListResponseBodyData {
	s.Language = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetReleaseTime(v string) *DescribeImageListResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetRenderingType(v string) *DescribeImageListResponseBodyData {
	s.RenderingType = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetStatus(v string) *DescribeImageListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetSystemType(v string) *DescribeImageListResponseBodyData {
	s.SystemType = &v
	return s
}

func (s *DescribeImageListResponseBodyData) Validate() error {
	if s.ImageBizTags != nil {
		for _, item := range s.ImageBizTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageListResponseBodyDataImageBizTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageListResponseBodyDataImageBizTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListResponseBodyDataImageBizTags) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponseBodyDataImageBizTags) GetKey() *string {
	return s.Key
}

func (s *DescribeImageListResponseBodyDataImageBizTags) GetValue() *string {
	return s.Value
}

func (s *DescribeImageListResponseBodyDataImageBizTags) SetKey(v string) *DescribeImageListResponseBodyDataImageBizTags {
	s.Key = &v
	return s
}

func (s *DescribeImageListResponseBodyDataImageBizTags) SetValue(v string) *DescribeImageListResponseBodyDataImageBizTags {
	s.Value = &v
	return s
}

func (s *DescribeImageListResponseBodyDataImageBizTags) Validate() error {
	return dara.Validate(s)
}
