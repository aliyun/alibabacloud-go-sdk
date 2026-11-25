// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiEngineEventStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAiEngineEventStatusResponseBody
	GetRequestId() *string
}

type ModifyAiEngineEventStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01A****E2-1219-019A-8B44-2557*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAiEngineEventStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiEngineEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAiEngineEventStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAiEngineEventStatusResponseBody) SetRequestId(v string) *ModifyAiEngineEventStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAiEngineEventStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
