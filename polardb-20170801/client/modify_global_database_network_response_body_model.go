// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
}

type ModifyGlobalDatabaseNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *ModifyGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
