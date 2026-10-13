// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowAutoRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelDataFlowAutoRefreshRequest
	GetClientToken() *string
	SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest
	GetFileSystemId() *string
	SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest
	GetRefreshPath() *string
}

type CancelDataFlowAutoRefreshRequest struct {
	// Ensures the idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique among different requests.
	//
	// The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring fees.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating the data flow. The dry run checks whether required parameters are specified, whether the request format is valid, and whether business dependencies are met. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but DataFlowId is empty.
	//
	// - false (default): Sends a normal request. After the check succeeds, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directory for which you want to cancel the automatic update configuration.
	//
	// Limits:
	//
	// - The length must be 2 to 1,024 characters.
	//
	// - The value must be encoded in UTF-8.
	//
	// - The value must start and end with a forward slash (/).
	//
	// > The directory must be an existing directory in CPFS and must be in a fileset for which data flow is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CancelDataFlowAutoRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelDataFlowAutoRefreshRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CancelDataFlowAutoRefreshRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelDataFlowAutoRefreshRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelDataFlowAutoRefreshRequest) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *CancelDataFlowAutoRefreshRequest) SetClientToken(v string) *CancelDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest {
	s.RefreshPath = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) Validate() error {
	return dara.Validate(s)
}
