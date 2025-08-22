// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaSpecificConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnIpaSpecificConfigResponseBody
	GetRequestId() *string
}

type DeleteDcdnIpaSpecificConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnIpaSpecificConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnIpaSpecificConfigResponseBody) SetRequestId(v string) *DeleteDcdnIpaSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
