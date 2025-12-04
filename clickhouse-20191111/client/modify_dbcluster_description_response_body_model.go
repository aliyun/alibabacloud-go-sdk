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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
