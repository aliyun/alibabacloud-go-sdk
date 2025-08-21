// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQpsModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQpsModeResponseBody
	GetRequestId() *string
}

type ModifyQpsModeResponseBody struct {
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 48859E14-A9FB-4100-99FF-AAB75CA46776
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQpsModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQpsModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQpsModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQpsModeResponseBody) SetRequestId(v string) *ModifyQpsModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQpsModeResponseBody) Validate() error {
	return dara.Validate(s)
}
