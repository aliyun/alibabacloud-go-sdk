// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyAppAgentResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyAppAgentResponseBody
	GetRequestId() *string
}

type ModifyAppAgentResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. It can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyAppAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppAgentResponseBody) SetData(v bool) *ModifyAppAgentResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAppAgentResponseBody) SetRequestId(v string) *ModifyAppAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
