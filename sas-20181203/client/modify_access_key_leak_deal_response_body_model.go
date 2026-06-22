// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessKeyLeakDealResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccessKeyLeakDealResponseBody
	GetRequestId() *string
}

type ModifyAccessKeyLeakDealResponseBody struct {
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// DD4617B4-133A-53C8-ADAE-7B30FF89****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessKeyLeakDealResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessKeyLeakDealResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessKeyLeakDealResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccessKeyLeakDealResponseBody) SetRequestId(v string) *ModifyAccessKeyLeakDealResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccessKeyLeakDealResponseBody) Validate() error {
	return dara.Validate(s)
}
