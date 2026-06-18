// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionDetailFileListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCommissionDetailFileListResponseBody
	GetCode() *string
	SetData(v *GetCommissionDetailFileListResponseBodyData) *GetCommissionDetailFileListResponseBody
	GetData() *GetCommissionDetailFileListResponseBodyData
	SetMessage(v string) *GetCommissionDetailFileListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCommissionDetailFileListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCommissionDetailFileListResponseBody
	GetSuccess() *bool
}

type GetCommissionDetailFileListResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetCommissionDetailFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCommissionDetailFileListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCommissionDetailFileListResponseBody) GetData() *GetCommissionDetailFileListResponseBodyData {
	return s.Data
}

func (s *GetCommissionDetailFileListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCommissionDetailFileListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommissionDetailFileListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCommissionDetailFileListResponseBody) SetCode(v string) *GetCommissionDetailFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetData(v *GetCommissionDetailFileListResponseBodyData) *GetCommissionDetailFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetMessage(v string) *GetCommissionDetailFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetRequestId(v string) *GetCommissionDetailFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetSuccess(v bool) *GetCommissionDetailFileListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCommissionDetailFileListResponseBodyData struct {
	// The billing month.
	//
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The file list.
	FileList []*GetCommissionDetailFileListResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// The UID of the partner.
	//
	// example:
	//
	// 1112332432
	PartnerUid *string `json:"PartnerUid,omitempty" xml:"PartnerUid,omitempty"`
}

func (s GetCommissionDetailFileListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBodyData) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetCommissionDetailFileListResponseBodyData) GetFileList() []*GetCommissionDetailFileListResponseBodyDataFileList {
	return s.FileList
}

func (s *GetCommissionDetailFileListResponseBodyData) GetPartnerUid() *string {
	return s.PartnerUid
}

func (s *GetCommissionDetailFileListResponseBodyData) SetBillMonth(v string) *GetCommissionDetailFileListResponseBodyData {
	s.BillMonth = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyData) SetFileList(v []*GetCommissionDetailFileListResponseBodyDataFileList) *GetCommissionDetailFileListResponseBodyData {
	s.FileList = v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyData) SetPartnerUid(v string) *GetCommissionDetailFileListResponseBodyData {
	s.PartnerUid = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyData) Validate() error {
	if s.FileList != nil {
		for _, item := range s.FileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCommissionDetailFileListResponseBodyDataFileList struct {
	// The push status of the OSS file. Valid values: Processing, Succeeded, or Failed.
	//
	// example:
	//
	// 处理中
	BucketSyncStatus *string `json:"BucketSyncStatus,omitempty" xml:"BucketSyncStatus,omitempty"`
	// The policy name.
	//
	// example:
	//
	// 参考【APS佣金查询下载功能】的政策名称
	CommissionPolicyName *string `json:"CommissionPolicyName,omitempty" xml:"CommissionPolicyName,omitempty"`
	// The file name.
	//
	// example:
	//
	// 佣金202502021112
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// example:
	//
	// 总代政策为拓渠、普通政策为拓客
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file URL.
	//
	// example:
	//
	// aps.ailyun.com/file/download?resourceId=1234&type=1
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetCommissionDetailFileListResponseBodyDataFileList) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBodyDataFileList) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) GetBucketSyncStatus() *string {
	return s.BucketSyncStatus
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) GetCommissionPolicyName() *string {
	return s.CommissionPolicyName
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) GetFileName() *string {
	return s.FileName
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) GetFileType() *string {
	return s.FileType
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetBucketSyncStatus(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.BucketSyncStatus = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetCommissionPolicyName(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.CommissionPolicyName = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetFileName(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.FileName = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetFileType(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.FileType = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetFileUrl(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.FileUrl = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) Validate() error {
	return dara.Validate(s)
}
