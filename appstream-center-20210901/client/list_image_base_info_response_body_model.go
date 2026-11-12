// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListImageBaseInfoResponseBody
	GetCode() *string
	SetCount(v int32) *ListImageBaseInfoResponseBody
	GetCount() *int32
	SetData(v []*ListImageBaseInfoResponseBodyData) *ListImageBaseInfoResponseBody
	GetData() []*ListImageBaseInfoResponseBodyData
	SetMessage(v string) *ListImageBaseInfoResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListImageBaseInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageBaseInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListImageBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageBaseInfoResponseBody
	GetSuccess() *bool
}

type ListImageBaseInfoResponseBody struct {
	// The result code of the request.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of images that meet the query conditions.
	//
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of basic image information on the current page.
	Data []*ListImageBaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The result message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success, and a value of false indicates failure.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageBaseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListImageBaseInfoResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListImageBaseInfoResponseBody) GetData() []*ListImageBaseInfoResponseBodyData {
	return s.Data
}

func (s *ListImageBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListImageBaseInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageBaseInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageBaseInfoResponseBody) SetCode(v string) *ListImageBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetCount(v int32) *ListImageBaseInfoResponseBody {
	s.Count = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetData(v []*ListImageBaseInfoResponseBodyData) *ListImageBaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetMessage(v string) *ListImageBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetPageNumber(v int32) *ListImageBaseInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetPageSize(v int32) *ListImageBaseInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetRequestId(v string) *ListImageBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) SetSuccess(v bool) *ListImageBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageBaseInfoResponseBody) Validate() error {
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

type ListImageBaseInfoResponseBodyData struct {
	// The ID of the Alibaba Cloud account to which the image belongs. A value of 0 indicates a system image.
	//
	// example:
	//
	// 0
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The business line code of the image, which corresponds to the BizType request parameter.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The distribution name of the image operating system.
	//
	// example:
	//
	// Ubuntu 22.04 64-bit
	Distro *string `json:"Distro,omitempty" xml:"Distro,omitempty"`
	// The time when the image was created. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the image was last updated. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2025-04-25 15:13:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image ID. The ID of a custom image is generated by the system.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image type. A value of System indicates a system image, User indicates a custom image, Shared indicates a shared image, and Community indicates a community image.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The ID of the latest sub-version of the image.
	//
	// example:
	//
	// iv-xxx
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// The ID of the sub-version currently used to read the main image information for this image.
	//
	// example:
	//
	// iv-xxxx
	OnlineVersionId *string `json:"OnlineVersionId,omitempty" xml:"OnlineVersionId,omitempty"`
	// The image operating system type. A value of Windows indicates the Windows operating system, Linux indicates the Linux operating system, Android indicates the Android operating system, and Unknown indicates an unknown operating system.
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image package type. A value of Ecs indicates an ECS image, Vhd indicates a VHD image, Container indicates a container image, Ecs_Container indicates an ECS and container composite image, and K8s_Node indicates a Kubernetes node image.
	//
	// example:
	//
	// Ecs
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The product type to which the image belongs. Examples: CloudDesktop (WUYING Workspace), CloudApp (WUYING Cloud App), and WuyingServer (WUYING Workstation).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The image protocol type. A value of ASP indicates the Alibaba Cloud self-developed ASP protocol, and HDX indicates the HDX protocol.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The image status. A value of INIT indicates initialization, CREATING indicates creating, AVAILABLE indicates available, and CREATE_FAILED indicates creation failed.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListImageBaseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListImageBaseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImageBaseInfoResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListImageBaseInfoResponseBodyData) GetBizType() *int32 {
	return s.BizType
}

func (s *ListImageBaseInfoResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListImageBaseInfoResponseBodyData) GetDistro() *string {
	return s.Distro
}

func (s *ListImageBaseInfoResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListImageBaseInfoResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListImageBaseInfoResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageBaseInfoResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *ListImageBaseInfoResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *ListImageBaseInfoResponseBodyData) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *ListImageBaseInfoResponseBodyData) GetOnlineVersionId() *string {
	return s.OnlineVersionId
}

func (s *ListImageBaseInfoResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *ListImageBaseInfoResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *ListImageBaseInfoResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *ListImageBaseInfoResponseBodyData) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListImageBaseInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListImageBaseInfoResponseBodyData) SetAliUid(v int64) *ListImageBaseInfoResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetBizType(v int32) *ListImageBaseInfoResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetDescription(v string) *ListImageBaseInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetDistro(v string) *ListImageBaseInfoResponseBodyData {
	s.Distro = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetGmtCreate(v string) *ListImageBaseInfoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetGmtModified(v string) *ListImageBaseInfoResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetImageId(v string) *ListImageBaseInfoResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetImageName(v string) *ListImageBaseInfoResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetImageType(v string) *ListImageBaseInfoResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetLatestVersionId(v string) *ListImageBaseInfoResponseBodyData {
	s.LatestVersionId = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetOnlineVersionId(v string) *ListImageBaseInfoResponseBodyData {
	s.OnlineVersionId = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetOsType(v string) *ListImageBaseInfoResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetPackageType(v string) *ListImageBaseInfoResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetProductType(v string) *ListImageBaseInfoResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetProtocolType(v string) *ListImageBaseInfoResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) SetStatus(v string) *ListImageBaseInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListImageBaseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
