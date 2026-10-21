// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToFilesystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachVscToFilesystemsRequest
	GetClientToken() *string
	SetResourceIds(v []*AttachVscToFilesystemsRequestResourceIds) *AttachVscToFilesystemsRequest
	GetResourceIds() []*AttachVscToFilesystemsRequestResourceIds
	SetRoleChain(v []*AttachVscToFilesystemsRequestRoleChain) *AttachVscToFilesystemsRequest
	GetRoleChain() []*AttachVscToFilesystemsRequestRoleChain
}

type AttachVscToFilesystemsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client and ensure that the value is unique across different requests.
	//
	// The client token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the file systems and virtual storage channels (VSCs). You can specify up to 10 IDs per batch. This parameter must be passed in a flattened format, such as ResourceIds.1.FileSystemId=bmcpfs-xxx&ResourceIds.1.VscId=vsc-xxx. JSON arrays are not supported.
	//
	// This parameter is required.
	ResourceIds []*AttachVscToFilesystemsRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The role chain.
	//
	// > Required only for cross-account scenarios.
	RoleChain []*AttachVscToFilesystemsRequestRoleChain `json:"RoleChain,omitempty" xml:"RoleChain,omitempty" type:"Repeated"`
}

func (s AttachVscToFilesystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsRequest) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachVscToFilesystemsRequest) GetResourceIds() []*AttachVscToFilesystemsRequestResourceIds {
	return s.ResourceIds
}

func (s *AttachVscToFilesystemsRequest) GetRoleChain() []*AttachVscToFilesystemsRequestRoleChain {
	return s.RoleChain
}

func (s *AttachVscToFilesystemsRequest) SetClientToken(v string) *AttachVscToFilesystemsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVscToFilesystemsRequest) SetResourceIds(v []*AttachVscToFilesystemsRequestResourceIds) *AttachVscToFilesystemsRequest {
	s.ResourceIds = v
	return s
}

func (s *AttachVscToFilesystemsRequest) SetRoleChain(v []*AttachVscToFilesystemsRequestRoleChain) *AttachVscToFilesystemsRequest {
	s.RoleChain = v
	return s
}

func (s *AttachVscToFilesystemsRequest) Validate() error {
	if s.ResourceIds != nil {
		for _, item := range s.ResourceIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoleChain != nil {
		for _, item := range s.RoleChain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachVscToFilesystemsRequestResourceIds struct {
	// The ID of the file system.
	//
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the VSC. You can call the ListVscs operation of the Lingjun AI Computing Service eflo-controller/2022-12-15 API to query the ID. If no VSC is created, you can call the CreateVsc operation to create one.
	//
	// example:
	//
	// vsc-8vb864o3ppwfvh****
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s AttachVscToFilesystemsRequestResourceIds) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsRequestResourceIds) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsRequestResourceIds) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AttachVscToFilesystemsRequestResourceIds) GetVscId() *string {
	return s.VscId
}

func (s *AttachVscToFilesystemsRequestResourceIds) SetFileSystemId(v string) *AttachVscToFilesystemsRequestResourceIds {
	s.FileSystemId = &v
	return s
}

func (s *AttachVscToFilesystemsRequestResourceIds) SetVscId(v string) *AttachVscToFilesystemsRequestResourceIds {
	s.VscId = &v
	return s
}

func (s *AttachVscToFilesystemsRequestResourceIds) Validate() error {
	return dara.Validate(s)
}

type AttachVscToFilesystemsRequestRoleChain struct {
	// The UID of the Alibaba Cloud account. This parameter specifies the Alibaba Cloud account for which the service assumes the role.
	//
	// example:
	//
	// 176498472570****
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// The resource descriptor of the specified role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// acs:ram::123472233814****:role/aliyunnasclientvsc****
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the role. Valid values: service and user.
	//
	// example:
	//
	// service
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AttachVscToFilesystemsRequestRoleChain) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsRequestRoleChain) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsRequestRoleChain) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *AttachVscToFilesystemsRequestRoleChain) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AttachVscToFilesystemsRequestRoleChain) GetRoleType() *string {
	return s.RoleType
}

func (s *AttachVscToFilesystemsRequestRoleChain) SetAssumeRoleFor(v string) *AttachVscToFilesystemsRequestRoleChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *AttachVscToFilesystemsRequestRoleChain) SetRoleArn(v string) *AttachVscToFilesystemsRequestRoleChain {
	s.RoleArn = &v
	return s
}

func (s *AttachVscToFilesystemsRequestRoleChain) SetRoleType(v string) *AttachVscToFilesystemsRequestRoleChain {
	s.RoleType = &v
	return s
}

func (s *AttachVscToFilesystemsRequestRoleChain) Validate() error {
	return dara.Validate(s)
}
