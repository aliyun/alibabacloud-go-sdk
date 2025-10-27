// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody
	GetRequestId() *string
}

type ModifyDBClusterDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17F57FEE-EA4F-4337-8D2E-9C23CAA63D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
