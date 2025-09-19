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
	// The ID of the request source. Default value: **aegis**. Valid values:
	//
	// 	- **aegis**: Server Guard
	//
	// 	- **sas**: Security Center
	//
	// >  If you use Server Guard, set the value to **aegis**. If you use Security Center, set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
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
