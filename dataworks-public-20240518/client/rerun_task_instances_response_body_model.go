// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RerunTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *RerunTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type RerunTaskInstancesResponseBody struct {
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the batch operation. The structure is a map in which the key is the node instance ID and the value is the result information.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s RerunTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerunTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RerunTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerunTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *RerunTaskInstancesResponseBody) SetRequestId(v string) *RerunTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *RerunTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *RerunTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
