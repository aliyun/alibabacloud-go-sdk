// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDcdnUserConfigResponseBody
	GetRequestId() *string
}

type SetDcdnUserConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnUserConfigResponseBody) SetRequestId(v string) *SetDcdnUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
