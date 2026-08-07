// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *AddAppAgentResponseBody
	GetData() *string
	SetRequestId(v string) *AddAppAgentResponseBody
	GetRequestId() *string
}

type AddAppAgentResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// ag.abcxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAppAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAppAgentResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppAgentResponseBody) GetData() *string {
	return s.Data
}

func (s *AddAppAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAppAgentResponseBody) SetData(v string) *AddAppAgentResponseBody {
	s.Data = &v
	return s
}

func (s *AddAppAgentResponseBody) SetRequestId(v string) *AddAppAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAppAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
