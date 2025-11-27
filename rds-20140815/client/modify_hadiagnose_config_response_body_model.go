// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHADiagnoseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHADiagnoseConfigResponseBody
	GetRequestId() *string
}

type ModifyHADiagnoseConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EFA6DF3-5247-4D9D-80AA-68765BE6D5EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHADiagnoseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHADiagnoseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHADiagnoseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHADiagnoseConfigResponseBody) SetRequestId(v string) *ModifyHADiagnoseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHADiagnoseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
