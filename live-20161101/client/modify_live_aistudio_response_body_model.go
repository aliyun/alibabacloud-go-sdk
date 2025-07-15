// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveAIStudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLiveAIStudioResponseBody
	GetRequestId() *string
}

type ModifyLiveAIStudioResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0BA6B3C4-525A-5381-A2B0-5351323F31C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveAIStudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveAIStudioResponseBody) SetRequestId(v string) *ModifyLiveAIStudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveAIStudioResponseBody) Validate() error {
	return dara.Validate(s)
}
