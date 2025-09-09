// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceADAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceADAuthServerResponseBody
	GetRequestId() *string
}

type ModifyInstanceADAuthServerResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C9E97677-BD74-584B-AFCE-948C2A70BB82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceADAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceADAuthServerResponseBody) SetRequestId(v string) *ModifyInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceADAuthServerResponseBody) Validate() error {
	return dara.Validate(s)
}
