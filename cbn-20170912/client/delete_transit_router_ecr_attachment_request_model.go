// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterEcrAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterEcrAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterEcrAttachmentRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTransitRouterEcrAttachmentRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterEcrAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterEcrAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterEcrAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterEcrAttachmentRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DeleteTransitRouterEcrAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteTransitRouterEcrAttachmentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, a request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcibly delete the ECR connection. Valid values:
	//
	// 	- **false*	- (default): checks for resource dependencies such as associated forwarding and route learning before deleting the ECR connection. If such resources exist, the ECR connection is not deleted and an error message is returned.
	//
	// 	- **true**: deletes the ECR connections and all dependent resources.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the ECR connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterEcrAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterEcrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterEcrAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterEcrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterEcrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetForce(v bool) *DeleteTransitRouterEcrAttachmentRequest {
	s.Force = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterEcrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterEcrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterEcrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterEcrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterEcrAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
