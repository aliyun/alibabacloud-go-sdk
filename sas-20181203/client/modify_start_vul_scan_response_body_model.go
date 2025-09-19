// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStartVulScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStartVulScanResponseBody
	GetRequestId() *string
}

type ModifyStartVulScanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4066CAC3-F83A-4729-9995-A5558A61B546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStartVulScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStartVulScanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStartVulScanResponseBody) SetRequestId(v string) *ModifyStartVulScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStartVulScanResponseBody) Validate() error {
	return dara.Validate(s)
}
