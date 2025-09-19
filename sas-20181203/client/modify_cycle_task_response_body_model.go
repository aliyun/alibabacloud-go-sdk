// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCycleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCycleTaskResponseBody
	GetRequestId() *string
}

type ModifyCycleTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AA33E30-7192-5648-93CD-D0E476A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCycleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCycleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCycleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCycleTaskResponseBody) SetRequestId(v string) *ModifyCycleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCycleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
