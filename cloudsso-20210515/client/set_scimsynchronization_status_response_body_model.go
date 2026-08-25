// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSCIMSynchronizationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSCIMSynchronizationStatusResponseBody
	GetRequestId() *string
}

type SetSCIMSynchronizationStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3BF1FC78-5D20-54CC-BAEB-8CC33AE21D01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSCIMSynchronizationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSCIMSynchronizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSCIMSynchronizationStatusResponseBody) SetRequestId(v string) *SetSCIMSynchronizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSCIMSynchronizationStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
