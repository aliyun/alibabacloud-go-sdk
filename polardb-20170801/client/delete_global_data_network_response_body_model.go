// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDataNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGlobalDataNetworkResponseBody
	GetRequestId() *string
}

type DeleteGlobalDataNetworkResponseBody struct {
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalDataNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDataNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDataNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalDataNetworkResponseBody) SetRequestId(v string) *DeleteGlobalDataNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalDataNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
