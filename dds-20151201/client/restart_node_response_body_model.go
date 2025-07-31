// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartNodeResponseBody
	GetRequestId() *string
}

type RestartNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ECBCA991-XXXX-XXXX-834C-B3E8007F33AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RestartNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartNodeResponseBody) SetRequestId(v string) *RestartNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
