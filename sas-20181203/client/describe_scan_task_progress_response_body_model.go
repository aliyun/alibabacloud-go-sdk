// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScanTaskProgressResponseBody
	GetRequestId() *string
	SetScanTaskProgress(v string) *DescribeScanTaskProgressResponseBody
	GetScanTaskProgress() *string
	SetTargetInfo(v string) *DescribeScanTaskProgressResponseBody
	GetTargetInfo() *string
}

type DescribeScanTaskProgressResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EA15BA8A-D631-4375-8D40-CB7C769B0279
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The progress of the virus scan task. Valid values:
	//
	// 	- **init**: The task is being initialized.
	//
	// 	- **Processing**: The task is running.
	//
	// 	- **Success**: The task is complete.
	//
	// 	- **Failed**: The task fails.
	//
	// example:
	//
	// Success
	ScanTaskProgress *string `json:"ScanTaskProgress,omitempty" xml:"ScanTaskProgress,omitempty"`
	// The information about the asset on which the virus scan task runs. The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **type**: the type of the asset on which you want to perform a virus scan task. Valid values:
	//
	//     	- **groupId**: server group.
	//
	//     	- **uuid**: server.
	//
	// 	- **name**: the name of the server group or server.
	//
	// 	- **target**: the asset on which the virus scan task runs. The value of this field varies based on the value of the type field.
	//
	//     	- If the **type*	- field is set to **groupId**, the value of this field is the ID of the server group.
	//
	//     	- If the **type*	- field is set to **uuid**, the value of this field is the universally unique identifier (UUID) of the server.
	//
	// example:
	//
	// [{"type":"uuid","name":"host001","target":"503201a7-14c6-4280-801b-1169ed42****"}]
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
}

func (s DescribeScanTaskProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScanTaskProgressResponseBody) GetScanTaskProgress() *string {
	return s.ScanTaskProgress
}

func (s *DescribeScanTaskProgressResponseBody) GetTargetInfo() *string {
	return s.TargetInfo
}

func (s *DescribeScanTaskProgressResponseBody) SetRequestId(v string) *DescribeScanTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetScanTaskProgress(v string) *DescribeScanTaskProgressResponseBody {
	s.ScanTaskProgress = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetTargetInfo(v string) *DescribeScanTaskProgressResponseBody {
	s.TargetInfo = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
