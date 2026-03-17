// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagHaResponseBody
	GetRequestId() *string
}

type ModifySagHaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3200E8A3-563F-4FFC-8BDB-0F1263FA69E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagHaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagHaResponseBody) SetRequestId(v string) *ModifySagHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagHaResponseBody) Validate() error {
	return dara.Validate(s)
}
