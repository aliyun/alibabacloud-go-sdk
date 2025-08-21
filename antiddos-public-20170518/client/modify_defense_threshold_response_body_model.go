// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseThresholdResponseBody
	GetRequestId() *string
}

type ModifyDefenseThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0607F8-A9F3-5E11-977B-D59CD58C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseThresholdResponseBody) SetRequestId(v string) *ModifyDefenseThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}
