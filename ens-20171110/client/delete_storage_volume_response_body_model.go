// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStorageVolumeResponseBody
	GetRequestId() *string
}

type DeleteStorageVolumeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 853D6E71-E087-1557-B65C-32BFBEE5CD97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStorageVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStorageVolumeResponseBody) SetRequestId(v string) *DeleteStorageVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}
