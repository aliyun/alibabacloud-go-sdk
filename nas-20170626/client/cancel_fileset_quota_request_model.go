// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFilesetQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelFilesetQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *CancelFilesetQuotaRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelFilesetQuotaRequest
	GetFileSystemId() *string
	SetFsetId(v string) *CancelFilesetQuotaRequest
	GetFsetId() *string
}

type CancelFilesetQuotaRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the token from your client and make sure that the token is unique across different requests.
	//
	// The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request.
	//
	// A dry run checks the validity of parameters and business constraints. It does not actually cancel the fileset quota or incur any fees.
	//
	// Valid values:
	//
	// - true: Sends a check request and does not cancel the quota. The check items include whether required parameters are specified, request format, and business limits. If the check fails, the corresponding error is returned. If the check succeeds, the HTTP status code 200 is returned.
	//
	// - false (default value): Sends a normal request and directly cancels the fileset quota after the check succeeds.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the CPFS for Lingjun file system. The ID must start with `bmcpfs-`, for example, bmcpfs-290w65p03ok64ya****. To use CPFS for Lingjun, you must first submit a ticket to apply for the service whitelist and then create the file system in the NAS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The fileset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s CancelFilesetQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelFilesetQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelFilesetQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelFilesetQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelFilesetQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelFilesetQuotaRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *CancelFilesetQuotaRequest) SetClientToken(v string) *CancelFilesetQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetDryRun(v bool) *CancelFilesetQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetFileSystemId(v string) *CancelFilesetQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetFsetId(v string) *CancelFilesetQuotaRequest {
	s.FsetId = &v
	return s
}

func (s *CancelFilesetQuotaRequest) Validate() error {
	return dara.Validate(s)
}
