// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevelopmentModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDevelopmentModeResponseBody
	GetRequestId() *string
}

type UpdateDevelopmentModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDevelopmentModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevelopmentModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevelopmentModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDevelopmentModeResponseBody) SetRequestId(v string) *UpdateDevelopmentModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevelopmentModeResponseBody) Validate() error {
	return dara.Validate(s)
}
