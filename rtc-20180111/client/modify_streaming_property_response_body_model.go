// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStreamingPropertyResponseBody
	GetRequestId() *string
}

type ModifyStreamingPropertyResponseBody struct {
	// example:
	//
	// E7C44674-9065-5BBA-AB77-A5F20908E73B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStreamingPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStreamingPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStreamingPropertyResponseBody) SetRequestId(v string) *ModifyStreamingPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStreamingPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
