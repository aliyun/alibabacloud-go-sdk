// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityEventMarkMissListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityEventMarkMissListResponseBody
	GetRequestId() *string
}

type DeleteSecurityEventMarkMissListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6BA51E2XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityEventMarkMissListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityEventMarkMissListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityEventMarkMissListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityEventMarkMissListResponseBody) SetRequestId(v string) *DeleteSecurityEventMarkMissListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityEventMarkMissListResponseBody) Validate() error {
	return dara.Validate(s)
}
