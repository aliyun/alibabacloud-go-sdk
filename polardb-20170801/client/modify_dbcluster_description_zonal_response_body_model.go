// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterDescriptionZonalResponseBody
	GetRequestId() *string
}

type ModifyDBClusterDescriptionZonalResponseBody struct {
	// example:
	//
	// 65D7ACE6-4A61-4B6E-B357-8CB24A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterDescriptionZonalResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
