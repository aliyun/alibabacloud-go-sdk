// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnFullDomainsBlockIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockInterval(v int32) *SetCdnFullDomainsBlockIPRequest
	GetBlockInterval() *int32
	SetIPList(v string) *SetCdnFullDomainsBlockIPRequest
	GetIPList() *string
	SetOperationType(v string) *SetCdnFullDomainsBlockIPRequest
	GetOperationType() *string
	SetUpdateType(v string) *SetCdnFullDomainsBlockIPRequest
	GetUpdateType() *string
}

type SetCdnFullDomainsBlockIPRequest struct {
	// The duration for which IP addresses or CIDR blocks are blocked. Unit: seconds. The value **0*	- specifies that IP addresses or CIDR blocks are permanently blocked. This parameter is available only if you set **OperationType*	- to **block**.
	//
	// example:
	//
	// 3000
	BlockInterval *int32 `json:"BlockInterval,omitempty" xml:"BlockInterval,omitempty"`
	// The IP addresses that you want to block or unblock. Separate multiple IP addresses with commas (,). You can specify up to 1,000 IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.XXX.XXX.1,2.XXX.XXX.2
	IPList *string `json:"IPList,omitempty" xml:"IPList,omitempty"`
	// The type of the operation.
	//
	// 	- block
	//
	// 	- unblock
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The type of the blocking duration. This parameter is available only if you set **OperationType*	- to **block**. Valid values:
	//
	// 	- **cover**: The blocking duration that is specified in the request takes effect.
	//
	// 	- **uncover**: The longer one of the blocking duration that is specified in the request and the remaining blocking duration takes effect.
	//
	// 	- Default value: cover.
	//
	// example:
	//
	// cover
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s SetCdnFullDomainsBlockIPRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCdnFullDomainsBlockIPRequest) GoString() string {
	return s.String()
}

func (s *SetCdnFullDomainsBlockIPRequest) GetBlockInterval() *int32 {
	return s.BlockInterval
}

func (s *SetCdnFullDomainsBlockIPRequest) GetIPList() *string {
	return s.IPList
}

func (s *SetCdnFullDomainsBlockIPRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetCdnFullDomainsBlockIPRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *SetCdnFullDomainsBlockIPRequest) SetBlockInterval(v int32) *SetCdnFullDomainsBlockIPRequest {
	s.BlockInterval = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPRequest) SetIPList(v string) *SetCdnFullDomainsBlockIPRequest {
	s.IPList = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPRequest) SetOperationType(v string) *SetCdnFullDomainsBlockIPRequest {
	s.OperationType = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPRequest) SetUpdateType(v string) *SetCdnFullDomainsBlockIPRequest {
	s.UpdateType = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPRequest) Validate() error {
	return dara.Validate(s)
}
