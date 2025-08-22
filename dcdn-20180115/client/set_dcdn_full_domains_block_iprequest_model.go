// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnFullDomainsBlockIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockInterval(v int32) *SetDcdnFullDomainsBlockIPRequest
	GetBlockInterval() *int32
	SetIPList(v string) *SetDcdnFullDomainsBlockIPRequest
	GetIPList() *string
	SetOperationType(v string) *SetDcdnFullDomainsBlockIPRequest
	GetOperationType() *string
	SetUpdateType(v string) *SetDcdnFullDomainsBlockIPRequest
	GetUpdateType() *string
}

type SetDcdnFullDomainsBlockIPRequest struct {
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
	// The action that you want to perform. Valid values:
	//
	// 	- **block**
	//
	// 	- **unblock**
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The type of the blocking duration. This parameter is available only if you set **OperationType*	- to **block**. Valid values:
	//
	// 	- **cover**: the blocking duration that is specified in the request takes effect.
	//
	// 	- **uncover**: the longer one of the blocking duration that is specified in the request and the remaining blocking duration takes effect.
	//
	// 	- Default value: cover.
	//
	// example:
	//
	// cover
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s SetDcdnFullDomainsBlockIPRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPRequest) GetBlockInterval() *int32 {
	return s.BlockInterval
}

func (s *SetDcdnFullDomainsBlockIPRequest) GetIPList() *string {
	return s.IPList
}

func (s *SetDcdnFullDomainsBlockIPRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetDcdnFullDomainsBlockIPRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetBlockInterval(v int32) *SetDcdnFullDomainsBlockIPRequest {
	s.BlockInterval = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetIPList(v string) *SetDcdnFullDomainsBlockIPRequest {
	s.IPList = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetOperationType(v string) *SetDcdnFullDomainsBlockIPRequest {
	s.OperationType = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetUpdateType(v string) *SetDcdnFullDomainsBlockIPRequest {
	s.UpdateType = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) Validate() error {
	return dara.Validate(s)
}
