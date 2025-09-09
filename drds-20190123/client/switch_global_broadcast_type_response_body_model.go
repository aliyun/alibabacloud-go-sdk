// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGlobalBroadcastTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SwitchGlobalBroadcastTypeResponseBody
	GetData() *bool
	SetRequestId(v string) *SwitchGlobalBroadcastTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchGlobalBroadcastTypeResponseBody
	GetSuccess() *bool
}

type SwitchGlobalBroadcastTypeResponseBody struct {
	// Indicates whether the mode of broadcast tables was switched from the multi-write mode to the asynchronous link mode.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGlobalBroadcastTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponseBody) GetData() *bool {
	return s.Data
}

func (s *SwitchGlobalBroadcastTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchGlobalBroadcastTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetData(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Data = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetRequestId(v string) *SwitchGlobalBroadcastTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetSuccess(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
