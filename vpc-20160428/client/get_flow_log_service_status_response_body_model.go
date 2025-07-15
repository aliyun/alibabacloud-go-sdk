// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowLogServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetFlowLogServiceStatusResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *GetFlowLogServiceStatusResponseBody
	GetRequestId() *string
}

type GetFlowLogServiceStatusResponseBody struct {
	// Indicates whether the flow log feature is enabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no You can call the [OpenFlowLogService](https://help.aliyun.com/document_detail/449637.html) operation to enable the flow log feature.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-06F83A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowLogServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowLogServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowLogServiceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetFlowLogServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowLogServiceStatusResponseBody) SetEnabled(v bool) *GetFlowLogServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetFlowLogServiceStatusResponseBody) SetRequestId(v string) *GetFlowLogServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowLogServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
