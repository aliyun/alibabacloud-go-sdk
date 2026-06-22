// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenLogShipperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *ModifyOpenLogShipperRequest
	GetFrom() *string
	SetResourceDirectoryAccountId(v int64) *ModifyOpenLogShipperRequest
	GetResourceDirectoryAccountId() *int64
}

type ModifyOpenLogShipperRequest struct {
	// The identifier of the request source. Default value: **aegis**. Valid values:
	//
	// - **aegis**: Server Guard edition.
	//
	// - **sas**: Security Center edition.
	//
	// > Server Guard users must use **aegis**, and Security Center users must use **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the member accounts in the resource directory (Alibaba Cloud account).
	//
	// > You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s ModifyOpenLogShipperRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenLogShipperRequest) GoString() string {
	return s.String()
}

func (s *ModifyOpenLogShipperRequest) GetFrom() *string {
	return s.From
}

func (s *ModifyOpenLogShipperRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyOpenLogShipperRequest) SetFrom(v string) *ModifyOpenLogShipperRequest {
	s.From = &v
	return s
}

func (s *ModifyOpenLogShipperRequest) SetResourceDirectoryAccountId(v int64) *ModifyOpenLogShipperRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyOpenLogShipperRequest) Validate() error {
	return dara.Validate(s)
}
