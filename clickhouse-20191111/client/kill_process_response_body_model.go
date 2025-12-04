// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *KillProcessResponseBody
	GetRequestId() *string
}

type KillProcessResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillProcessResponseBody) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillProcessResponseBody) SetRequestId(v string) *KillProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
