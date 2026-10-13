// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *StartDataFlowRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *StartDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *StartDataFlowRequest
	GetFileSystemId() *string
}

type StartDataFlowRequest struct {
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
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
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring fees.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating an instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly created.
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

func (s StartDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StartDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *StartDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StartDataFlowRequest) SetClientToken(v string) *StartDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDataFlowRequest) SetDataFlowId(v string) *StartDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StartDataFlowRequest) SetDryRun(v bool) *StartDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StartDataFlowRequest) SetFileSystemId(v string) *StartDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *StartDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
