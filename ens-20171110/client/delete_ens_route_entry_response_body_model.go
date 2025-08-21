// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnsRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteEnsRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnsRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnsRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnsRouteEntryResponseBody) SetRequestId(v string) *DeleteEnsRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnsRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
