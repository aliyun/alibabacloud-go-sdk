// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteProtocolServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteProtocolServiceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteProtocolServiceRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *DeleteProtocolServiceRequest
	GetProtocolServiceId() *string
}

type DeleteProtocolServiceRequest struct {
	// Ensures the idempotency of the request. Generate a value from your client that is unique across different requests.
	//
	// The ClientToken value supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary across API requests.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this request. A dry run checks parameter validity and dependencies without actually modifying the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without modifying the protocol service. The check items include whether required parameters are specified, the request format, and business limit dependencies. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the protocol service is directly modified.
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
	// The protocol service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteProtocolServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteProtocolServiceRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DeleteProtocolServiceRequest) SetClientToken(v string) *DeleteProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetDryRun(v bool) *DeleteProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetFileSystemId(v string) *DeleteProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetProtocolServiceId(v string) *DeleteProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
