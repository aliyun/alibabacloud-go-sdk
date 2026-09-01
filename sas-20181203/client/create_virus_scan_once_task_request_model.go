// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVirusScanOnceTaskRequest
	GetClientToken() *string
	SetParam(v string) *CreateVirusScanOnceTaskRequest
	GetParam() *string
	SetScanPath(v []*string) *CreateVirusScanOnceTaskRequest
	GetScanPath() []*string
	SetScanType(v string) *CreateVirusScanOnceTaskRequest
	GetScanType() *string
	SetSelectionKey(v string) *CreateVirusScanOnceTaskRequest
	GetSelectionKey() *string
}

type CreateVirusScanOnceTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The extended information field:
	//
	// - **additionType**: the extended scan type
	//
	// example:
	//
	// {\\"additionType\\":[\\"SCAN_MEMORY\\"]}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The scan path information to be transmitted if the scan type is custom scan.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The scan type of the virus scan. Valid values:
	//
	// - **system**: system scan
	//
	// - **user**: custom scan.
	//
	// example:
	//
	// system
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The key that stores asset information.
	//
	// > You can call the [GetAssetSelectionConfig](~~GetAssetSelectionConfig~~) operation to obtain this parameter.
	//
	// example:
	//
	// 845de1ec-4b08-42e1-b564-31321e48xxxx
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
}

func (s CreateVirusScanOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVirusScanOnceTaskRequest) GetParam() *string {
	return s.Param
}

func (s *CreateVirusScanOnceTaskRequest) GetScanPath() []*string {
	return s.ScanPath
}

func (s *CreateVirusScanOnceTaskRequest) GetScanType() *string {
	return s.ScanType
}

func (s *CreateVirusScanOnceTaskRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *CreateVirusScanOnceTaskRequest) SetClientToken(v string) *CreateVirusScanOnceTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetParam(v string) *CreateVirusScanOnceTaskRequest {
	s.Param = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetScanPath(v []*string) *CreateVirusScanOnceTaskRequest {
	s.ScanPath = v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetScanType(v string) *CreateVirusScanOnceTaskRequest {
	s.ScanType = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetSelectionKey(v string) *CreateVirusScanOnceTaskRequest {
	s.SelectionKey = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
