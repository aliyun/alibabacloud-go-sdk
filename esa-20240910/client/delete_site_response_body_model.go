// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSiteResponseBody
	GetRequestId() *string
}

type DeleteSiteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSiteResponseBody) SetRequestId(v string) *DeleteSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
