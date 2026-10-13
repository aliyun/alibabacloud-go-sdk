// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *DeleteDataFlowRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *DeleteDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteDataFlowRequest
	GetFileSystemId() *string
}

type DeleteDataFlowRequest struct {
	// Ensures the idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dataflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring fees.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating an instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request and creates the instance after the check passes.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *DeleteDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteDataFlowRequest) SetClientToken(v string) *DeleteDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDataFlowId(v string) *DeleteDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDryRun(v bool) *DeleteDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDataFlowRequest) SetFileSystemId(v string) *DeleteDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
