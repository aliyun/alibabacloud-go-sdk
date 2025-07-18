// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataRedistributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseDataRedistributeResponseBody
	GetRequestId() *string
}

type PauseDataRedistributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseDataRedistributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseDataRedistributeResponseBody) GoString() string {
	return s.String()
}

func (s *PauseDataRedistributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseDataRedistributeResponseBody) SetRequestId(v string) *PauseDataRedistributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseDataRedistributeResponseBody) Validate() error {
	return dara.Validate(s)
}
