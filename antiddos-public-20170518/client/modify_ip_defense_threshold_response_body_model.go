// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpDefenseThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpDefenseThresholdResponseBody
	GetRequestId() *string
}

type ModifyIpDefenseThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0607F8-A9F3-5E11-977B-D59CD58C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpDefenseThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpDefenseThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpDefenseThresholdResponseBody) SetRequestId(v string) *ModifyIpDefenseThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpDefenseThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}
