// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClickhouseEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClickhouseEngineResponseBody
	GetRequestId() *string
}

type ModifyClickhouseEngineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D761DA51-12F8-5457-AAA9-F52B9F436D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClickhouseEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClickhouseEngineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClickhouseEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClickhouseEngineResponseBody) SetRequestId(v string) *ModifyClickhouseEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClickhouseEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
