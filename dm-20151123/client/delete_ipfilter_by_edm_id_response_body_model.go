// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpfilterByEdmIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpfilterByEdmIdResponseBody
	GetRequestId() *string
}

type DeleteIpfilterByEdmIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// E3DFF97B-00CF-5333-8125-3D6819471984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpfilterByEdmIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpfilterByEdmIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpfilterByEdmIdResponseBody) SetRequestId(v string) *DeleteIpfilterByEdmIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpfilterByEdmIdResponseBody) Validate() error {
	return dara.Validate(s)
}
