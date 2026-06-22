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
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// EA15BA8A-D631-4375-8D40-CB7C769B0279
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The progress of the virus scan task. Valid values:
	//
	// - **init**: The scan task is being initialized.
	//
	// - **Processing**: The scan task is in progress.
	//
	// - **Success**: The scan task is complete.
	//
	// - **Failed**: The scan task failed.
	//
	// example:
	//
	// Success
	ScanTaskProgress *string `json:"ScanTaskProgress,omitempty" xml:"ScanTaskProgress,omitempty"`
	// The asset information scanned by the virus scan node. This parameter is a string converted from a JSON array in character format. The following fields are included:
	//
	// - **type**: The Asset Type on which the virus scan is executed. Valid values:
	//
	//     - **groupId**: server group.
	//
	//     - **uuid**: server.
	//
	// - **name**: The name of the server group or server.
	//
	// - **target**: The asset on which the virus scan is executed. The following describes the values of this field:
	//
	//     - If **type*	- is set to **groupId**, this field specifies the server group ID.
	//
	//     - If **type*	- is set to **uuid**, this field specifies the UUID of the server.
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
