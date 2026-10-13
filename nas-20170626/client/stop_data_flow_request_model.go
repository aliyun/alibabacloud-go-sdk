// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *StopDataFlowRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *StopDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *StopDataFlowRequest
	GetFileSystemId() *string
}

type StopDataFlowRequest struct {
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
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
	// A dry run checks parameter validity and inventory without actually creating an instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating an instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is created.
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
}

func (s StopDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StopDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *StopDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StopDataFlowRequest) SetClientToken(v string) *StopDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDataFlowRequest) SetDataFlowId(v string) *StopDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StopDataFlowRequest) SetDryRun(v bool) *StopDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StopDataFlowRequest) SetFileSystemId(v string) *StopDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *StopDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
