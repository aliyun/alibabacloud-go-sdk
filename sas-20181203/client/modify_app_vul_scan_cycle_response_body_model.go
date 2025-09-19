// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppVulScanCycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppVulScanCycleResponseBody
	GetRequestId() *string
}

type ModifyAppVulScanCycleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFEDC54D-70A2-5E56-A69B-E3D8AA8A5197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppVulScanCycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppVulScanCycleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppVulScanCycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppVulScanCycleResponseBody) SetRequestId(v string) *ModifyAppVulScanCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppVulScanCycleResponseBody) Validate() error {
	return dara.Validate(s)
}
