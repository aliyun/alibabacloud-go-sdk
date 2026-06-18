// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlCommissionDetailFileListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIntlCommissionDetailFileListResponseBody
	GetCode() *string
	SetData(v *GetIntlCommissionDetailFileListResponseBodyData) *GetIntlCommissionDetailFileListResponseBody
	GetData() *GetIntlCommissionDetailFileListResponseBodyData
	SetMessage(v string) *GetIntlCommissionDetailFileListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIntlCommissionDetailFileListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIntlCommissionDetailFileListResponseBody
	GetSuccess() *bool
}

type GetIntlCommissionDetailFileListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *GetIntlCommissionDetailFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message information.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIntlCommissionDetailFileListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntlCommissionDetailFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntlCommissionDetailFileListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIntlCommissionDetailFileListResponseBody) GetData() *GetIntlCommissionDetailFileListResponseBodyData {
	return s.Data
}

func (s *GetIntlCommissionDetailFileListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIntlCommissionDetailFileListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntlCommissionDetailFileListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIntlCommissionDetailFileListResponseBody) SetCode(v string) *GetIntlCommissionDetailFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBody) SetData(v *GetIntlCommissionDetailFileListResponseBodyData) *GetIntlCommissionDetailFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBody) SetMessage(v string) *GetIntlCommissionDetailFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBody) SetRequestId(v string) *GetIntlCommissionDetailFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBody) SetSuccess(v bool) *GetIntlCommissionDetailFileListResponseBody {
	s.Success = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIntlCommissionDetailFileListResponseBodyData struct {
	// The billing month.
	//
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The file list object.
	FileList []*GetIntlCommissionDetailFileListResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// The UID of the partner.
	//
	// example:
	//
	// 1112332432
	PartnerUid *string `json:"PartnerUid,omitempty" xml:"PartnerUid,omitempty"`
}

func (s GetIntlCommissionDetailFileListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIntlCommissionDetailFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) GetFileList() []*GetIntlCommissionDetailFileListResponseBodyDataFileList {
	return s.FileList
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) GetPartnerUid() *string {
	return s.PartnerUid
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) SetBillMonth(v string) *GetIntlCommissionDetailFileListResponseBodyData {
	s.BillMonth = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) SetFileList(v []*GetIntlCommissionDetailFileListResponseBodyDataFileList) *GetIntlCommissionDetailFileListResponseBodyData {
	s.FileList = v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) SetPartnerUid(v string) *GetIntlCommissionDetailFileListResponseBodyData {
	s.PartnerUid = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyData) Validate() error {
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

type GetIntlCommissionDetailFileListResponseBodyDataFileList struct {
	// The push status of the OSS file. Valid values:
	//
	// - 初始化状态: initialization status
	//
	// - 处理中: processing
	//
	// - 处理成功: processing succeeded
	//
	// - 处理失败: processing failed.
	//
	// example:
	//
	// 处理中
	BucketSyncStatus *string `json:"BucketSyncStatus,omitempty" xml:"BucketSyncStatus,omitempty"`
	// The commission policy name.
	//
	// example:
	//
	// 参考【APS佣金查询下载功能】的政策名称
	CommissionPolicyName *string `json:"CommissionPolicyName,omitempty" xml:"CommissionPolicyName,omitempty"`
	// The file name.
	//
	// example:
	//
	// 202606_FY27_TEST_4397912340.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetIntlCommissionDetailFileListResponseBodyDataFileList) String() string {
	return dara.Prettify(s)
}

func (s GetIntlCommissionDetailFileListResponseBodyDataFileList) GoString() string {
	return s.String()
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) GetBucketSyncStatus() *string {
	return s.BucketSyncStatus
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) GetCommissionPolicyName() *string {
	return s.CommissionPolicyName
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) GetFileName() *string {
	return s.FileName
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) SetBucketSyncStatus(v string) *GetIntlCommissionDetailFileListResponseBodyDataFileList {
	s.BucketSyncStatus = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) SetCommissionPolicyName(v string) *GetIntlCommissionDetailFileListResponseBodyDataFileList {
	s.CommissionPolicyName = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) SetFileName(v string) *GetIntlCommissionDetailFileListResponseBodyDataFileList {
	s.FileName = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponseBodyDataFileList) Validate() error {
	return dara.Validate(s)
}
