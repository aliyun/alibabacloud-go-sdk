// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
}

type DeleteGlobalDatabaseNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *DeleteGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
