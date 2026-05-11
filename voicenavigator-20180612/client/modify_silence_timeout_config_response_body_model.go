// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySilenceTimeoutConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySilenceTimeoutConfigResponseBody
	GetRequestId() *string
}

type ModifySilenceTimeoutConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySilenceTimeoutConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySilenceTimeoutConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySilenceTimeoutConfigResponseBody) SetRequestId(v string) *ModifySilenceTimeoutConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySilenceTimeoutConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
