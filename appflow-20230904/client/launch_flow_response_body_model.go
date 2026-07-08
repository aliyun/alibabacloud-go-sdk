// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *LaunchFlowResponseBody
	GetData() *string
	SetRequestId(v string) *LaunchFlowResponseBody
	GetRequestId() *string
}

type LaunchFlowResponseBody struct {
	// The returned object.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 965C0149-B68B-5A0C-AEA8-48DAEC7779BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LaunchFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LaunchFlowResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *LaunchFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LaunchFlowResponseBody) SetData(v string) *LaunchFlowResponseBody {
	s.Data = &v
	return s
}

func (s *LaunchFlowResponseBody) SetRequestId(v string) *LaunchFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
