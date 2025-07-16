// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFeatureViewResponseBody
	GetRequestId() *string
}

type DeleteFeatureViewResponseBody struct {
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFeatureViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFeatureViewResponseBody) SetRequestId(v string) *DeleteFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFeatureViewResponseBody) Validate() error {
	return dara.Validate(s)
}
