// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigInXMLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterConfigInXMLResponseBody
	GetRequestId() *string
}

type ModifyDBClusterConfigInXMLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BDD29EB1-BE76-5CFA-9068-D34B696310BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigInXMLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterConfigInXMLResponseBody) SetRequestId(v string) *ModifyDBClusterConfigInXMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponseBody) Validate() error {
	return dara.Validate(s)
}
